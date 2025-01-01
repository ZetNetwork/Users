package logrus_logger

import (
	"github.com/ZetNetwork/Users/pkg/logger"
	"log/slog"
	"os"
)

type loggerClient struct {
	client *slog.Logger
}

func NewLoggerClient() logger.ILogger {
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	)

	lC := loggerClient{
		client: slog.New(handler),
	}

	return lC
}

func (l loggerClient) Info(msg string, args ...any) {
	l.client.Info(msg, args...)
}

func (l loggerClient) Debug(msg string, args ...any) {
	l.client.Debug(msg, args...)
}

func (l loggerClient) Error(msg string, args ...any) {
	l.client.Error(msg, args...)
}
