package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bulbosaur/game-of-life/http/server"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
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

	ctx, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-c
	cancel()
	//  Завершим работу сервера
	shutDownFunc(ctx)

	return 0

}

// Настройки логгера
func setupLogger() *zap.Logger {
	// Настройка конфигурации логгера
	config := zap.NewProductionConfig()

	// Уровень логирования
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	// Настройка логгера с конфигурацией
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Ошибка настройки логгера: %v\n", err)
	}

	return logger
}
