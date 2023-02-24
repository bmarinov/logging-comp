package deepdown

import (
	"context"
	plainlog "log"
	"team2/topics/logging-comp/internal/logctx"

	"github.com/rs/zerolog/log"
)

func (r *SpicyBusinessLogic) ContextAsAService(ctx context.Context) {
	plainlog.Println("--- Logger in context:")

	// with context:
	sublogger := logctx.MustGetLoggerOrElse(ctx)
	sublogger.Warn().
		Msg("I am not a service locator I swear")
}

func (r *SpicyBusinessLogic) WithZeroSugar(ctx context.Context) {
	plainlog.Println("--- logger in context with zerolog helpers:")

	log.Ctx(ctx).Info().Msg("or like that")
}
