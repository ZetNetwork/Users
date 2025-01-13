package app

import (
	"github.com/ZetNetwork/Protos/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ZetNetwork/Users/internal/adapters/auth_adapter"
	"github.com/ZetNetwork/Users/internal/adapters/repository"
	"github.com/ZetNetwork/Users/internal/adapters/user_adapter"
	"github.com/ZetNetwork/Users/internal/domain/ports"
	"github.com/ZetNetwork/Users/internal/domain/usecase"
	"github.com/ZetNetwork/Users/internal/infrastructure/auth_server"
	"github.com/ZetNetwork/Users/internal/infrastructure/database"
	"github.com/ZetNetwork/Users/internal/infrastructure/database/go_postgres"
	"github.com/ZetNetwork/Users/internal/infrastructure/user_server"
)

type appProvider struct {
	pgConfig   database.IPGConfig
	authConfig auth_server.IAuthClientConfig
	userConfig user_server.IUserServerConfig

	pgClient *go_postgres.PostgreClient

	uc         usecase.IUserUseCase
	userServer *user_adapter.UsersServer

	auth ports.IAuthClient
	repo ports.IUserRepository
}

func newAppProvide() *appProvider {
	return &appProvider{}
}

func (a *appProvider) PGConfig() database.IPGConfig {
	if a.pgConfig == nil {
		config, err := database.NewPGConfig()
		if err != nil {
			return nil
		}
		a.pgConfig = config
	}
	return a.pgConfig
}

func (a *appProvider) AuthConfig() auth_server.IAuthClientConfig {
	if a.authConfig == nil {
		config, err := auth_server.NewAuthClientConfig()
		if err != nil {
			return nil
		}
		a.authConfig = config
	}
	return a.authConfig
}

func (a *appProvider) UserConfig() user_server.IUserServerConfig {
	if a.userConfig == nil {
		config, err := user_server.NewUserServerConfig()
		if err != nil {
			return nil
		}
		a.userConfig = config
	}
	return a.userConfig
}

func (a *appProvider) PGClient() *go_postgres.PostgreClient {
	if a.pgClient == nil {
		client, err := go_postgres.NewPostgresClient(a.PGConfig())
		if err != nil {
			return nil
		}
		a.pgClient = client
	}
	return a.pgClient
}

func (a *appProvider) UserUseCase() usecase.IUserUseCase {
	if a.uc == nil {
		uc := usecase.NewUserUseCase(a.Repo(), a.AuthClient())
		a.uc = uc
	}
	return a.uc
}

func (a *appProvider) Repo() ports.IUserRepository {
	if a.repo == nil {
		repo := repository.NewUserRepository(a.PGClient())
		a.repo = repo
	}
	return a.repo
}

func (a *appProvider) UserServer() *user_adapter.UsersServer {
	if a.userServer == nil {
		server := user_adapter.NewUserServer(a.UserUseCase())
		a.userServer = server
	}
	return a.userServer
}

func (a *appProvider) AuthClient() ports.IAuthClient {
	if a.auth == nil {
		conn, err := grpc.NewClient(a.AuthConfig().Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err.Error())
		}

		client := auth_v1.NewAuthV1Client(conn)
		auth := auth_adapter.NewAuthClient(client)
		a.auth = auth
	}
	return a.auth
}
