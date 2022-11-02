package zerolog

import (
	"context"

	"github.com/Thiht/fluor"
	"github.com/rs/zerolog"
)

func ZerologLogger(ctx context.Context, level fluor.LogLevel, err *fluor.FluentError) *fluor.FluentError {
	parsedLevel, parseErr := zerolog.ParseLevel(string(level))
	if parseErr != nil {
		parsedLevel = zerolog.ErrorLevel
	}
	zerolog.Ctx(ctx).WithLevel(parsedLevel).Err(err.Unwrap()).Fields(err.Fields()).Msg(err.Message())
	return err
}
