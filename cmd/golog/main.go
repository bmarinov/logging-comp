package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"team2/topics/logging-comp/cmd/golog/deepdown"
	"team2/topics/logging-comp/internal/logctx"
	"team2/topics/logging-comp/internal/logger"

	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("lets get loggin")
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	logger.Init()

	controller := deepdown.NewSpicyBusinessLogic(logger.GlobalLoggy)

	// "request" scope from here:
	requestContext := contextWithTraceIdEtc(ctx)

	handleRequest(requestContext, controller)
}

func handleRequest(requestContext context.Context, controller *deepdown.SpicyBusinessLogic) {

	first(controller, requestContext, logger.GlobalLoggy)
	second(controller, requestContext)
	third(controller, requestContext)
	fourth(controller, requestContext)
}

func first(foo *deepdown.SpicyBusinessLogic, requestContext context.Context, loggerInstance *zerolog.Logger) {
	// explicit dependency
	foo.LoggerDependency(requestContext)
	foo.AsFnArgument(requestContext, loggerInstance)
}

func second(foo *deepdown.SpicyBusinessLogic, requestContext context.Context) {
	// global logger
	foo.GlobalState(requestContext)
}

func third(foo *deepdown.SpicyBusinessLogic, ctx context.Context) {
	// a logger magically appears
	contextScopedLogger := logger.GlobalLoggy.With().Logger()
	poopyContext := logctx.LoadUp(ctx, contextScopedLogger)
	foo.ContextAsAService(poopyContext)

	// Zerolog API trying to hide the ugly bits:
	zerlogCtx := logger.GlobalLoggy.With().Str("compo_key_foo", "main_fn").Logger().WithContext(ctx)
	foo.WithZeroSugar(zerlogCtx)
}

func fourth(foo *deepdown.SpicyBusinessLogic, ctx context.Context) {
	foo.LittleHelpy(ctx)
}
