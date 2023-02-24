package logger

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GlobalLoggy is the global logger instance. Hopefully not nil
var GlobalLoggy *zerolog.Logger

func Init() {
	// configure
	zerolog.TimeFieldFormat = time.RFC3339
	newLogger := NewLoggy()

	GlobalLoggy = newLogger
	log.Logger = *newLogger
}

// NewLoggy returns a new sublogger preconfigured with global values from config.
func NewLoggy() *zerolog.Logger {
	// from env / config
	componentName := "loggy_testapp"

	logger := log.With().
		// static & global values
		Str("component", componentName).
		Logger()

	return &logger
}

// GetLoggyWithContext returns a sublogger with values from the context
func GetLoggyWithContext(ctx context.Context, logger *zerolog.Logger) zerolog.Logger {
	// extract from context

	correlationId := ctx.Value("correlationId")

	return logger.With().Str("correlationId", correlationId.(string)).
		Logger()
}

// ForContext creates a logger with values from the context.
func ForContext(ctx context.Context) zerolog.Logger {
	logmap := ctx.Value("magic_key")

	if logmap == nil {
		return *GlobalLoggy
	}
	mappy, ok := logmap.(map[string]interface{})
	if !ok {
		return *GlobalLoggy
	}

	sublogger := GlobalLoggy.With().Fields(mappy)

	return sublogger.Logger()
}
