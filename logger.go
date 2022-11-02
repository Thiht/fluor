package fluor

import (
	"context"
	"log"
)

type LogLevel string

const (
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

type LoggerFunc func(ctx context.Context, level LogLevel, err *FluentError) *FluentError

// Logger can be overridden to implement custom logic
var Logger LoggerFunc = defaultLogger

func defaultLogger(ctx context.Context, level LogLevel, err *FluentError) *FluentError {
	log.Printf("[%s] %v", level, err.Error())
	return err
}
