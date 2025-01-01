package logger

import (
	"context"
	"fmt"
)

type ctxLogger struct{}

func ContextWithLogger(ctx context.Context, logger ILogger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, logger)
}

func LoggerFromContext(ctx context.Context) ILogger {
	if logger, ok := ctx.Value(ctxLogger{}).(ILogger); ok {
		return logger
	}
	fmt.Println("Logger not found in context")
	return nil
}
