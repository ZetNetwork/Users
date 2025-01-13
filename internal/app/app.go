package app

import (
	"context"
	"fmt"
	"net"

	"github.com/ZetNetwork/Protos/pkg/users_v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/ZetNetwork/Users/internal/migrations"
	"github.com/ZetNetwork/Users/pkg/logger"
)

const (
	_EnvPath = ".env"
)

type initFunc func(ctx context.Context) error

type App struct {
	provider *appProvider
	server   *grpc.Server
}

func NewApp(ctx context.Context) *App {
	app := &App{}
	if err := app.InitDeps(ctx); err != nil {
		panic(err)
	}
	return app
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []initFunc{
		a.initConfig,
		a.initProvider,
		a.initServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	if err := migrations.MigrateDB(a.provider.PGClient()); err != nil {
		return fmt.Errorf("migrate static db: %v", err)
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	logger.LoggerFromContext(ctx).
		Debug("Loading env file")
	err := godotenv.Load(_EnvPath)
	if err != nil {
		return err
	}
	logger.LoggerFromContext(ctx).
		Info("Env file loaded")
	return nil
}

func (a *App) initProvider(ctx context.Context) error {
	logger.LoggerFromContext(ctx).
		Debug("Initializing app provider")
	a.provider = newAppProvide()
	logger.LoggerFromContext(ctx).
		Info("App provider initialized")
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	server := grpc.NewServer()
	users_v1.RegisterUsersV1Server(server, a.provider.UserServer())
	a.server = server
	return nil
}

func (a *App) Run(ctx context.Context) error {
	list, err := net.Listen("tcp", a.provider.UserConfig().Address())
	if err != nil {
		panic(err)
	}

	err = a.server.Serve(list)
	if err != nil {
		panic(err)
	}

	return nil
}
