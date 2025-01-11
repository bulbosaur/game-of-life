package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bulbosaur/game-of-life/config"
	"github.com/bulbosaur/game-of-life/http/server"
)

type Application struct {
	Cfg *config.Config
}

func New(config *config.Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) int {
	logger := setupLogger()

	shutDownFunc, err := server.Run(ctx, logger, a.Cfg.Height, a.Cfg.Width)
	if err != nil {
		logger.Error(err.Error())

		return 1
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	shutDownFunc(ctx)

	return 0

}

func setupLogger() *zap.Logger {
	config := zap.NewProductionConfig()

	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Logger error: %v\n", err)
	}

	return logger
}
