package main

import (
	"context"

	"github.com/ZetNetwork/Users/internal/app"
	"github.com/ZetNetwork/Users/pkg/logger"
	logrus_logger "github.com/ZetNetwork/Users/pkg/logger/logrus-logger"
)

func main() {
	ctx := logger.ContextWithLogger(context.Background(), logrus_logger.NewLoggerClient())
	root := app.NewApp(ctx)
	err := root.Run(ctx)
	if err != nil {
		return
	}
}
