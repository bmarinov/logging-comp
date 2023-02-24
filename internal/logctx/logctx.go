package logctx

import (
	"context"

	"github.com/rs/zerolog"
)

// MustGetLoggerOrElse is totally not a service locator.
func MustGetLoggerOrElse(ctx context.Context) zerolog.Logger {
	logger := ctx.Value(ctxKey)
	if logger == nil {
		panic("oh no someone messed up :)")
	}
	result, ok := logger.(zerolog.Logger)
	if !ok {
		panic("oh no cant cast value to Logger")
	}
	return result.With().Logger()
}

// ctxKey is the ctxKey for the logger in the context.
var ctxKey loggerContextKey

type loggerContextKey struct{}

// LoadUp adds the logger to the provided context with an appropriate key
func LoadUp(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, logger)
}
