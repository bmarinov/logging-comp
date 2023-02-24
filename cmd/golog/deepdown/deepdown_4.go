package deepdown

import (
	"context"
	plainlog "log"
	"team2/topics/logging-comp/internal/logger"
)

func (r SpicyBusinessLogic) LittleHelpy(ctx context.Context) {
	plainlog.Println("--- global loggy can be nice too:")

	sublogger := logger.ForContext(ctx)
	sublogger.Warn().
		Msg("where did all these values come from :o")
}
