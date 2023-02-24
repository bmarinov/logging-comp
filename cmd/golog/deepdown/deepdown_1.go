package deepdown

import (
	"context"
	plainlog "log"
	"team2/topics/logging-comp/internal/logger"

	"github.com/rs/zerolog"
)

func (r *SpicyBusinessLogic) LoggerDependency(ctx context.Context) {
	plainlog.Println("--- Logging from logger in struct:")

	// with context-aware logger:
	sublogger := logger.GetLoggyWithContext(ctx, &r.logger)
	sublogger.Info().Msg("Now with context")
}

func (r *SpicyBusinessLogic) AsFnArgument(ctx context.Context, loggy *zerolog.Logger) {
	plainlog.Println("--- Logger passed as arg:")

	// with context:
	sublogger := logger.GetLoggyWithContext(ctx, loggy)
	canAlsoLogHere(sublogger)
}

func canAlsoLogHere(loggy zerolog.Logger) {
	loggy.Info().Msg("^^^ yep, works from a function too")
}
