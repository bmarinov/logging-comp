package deepdown

import (
	"context"
	plainlog "log"
	"team2/topics/logging-comp/internal/logger"
)

func (r *SpicyBusinessLogic) GlobalState(ctx context.Context) {
	plainlog.Println("--- Logger from global state:")

	// with context:
	sublogger := logger.GetLoggyWithContext(ctx, logger.GlobalLoggy)
	sublogger.Warn().Msg("huhu KEINE PANIC! :)")
}
