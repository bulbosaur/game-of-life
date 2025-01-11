package main

import (
	"context"
	"log"
	"os"

	"github.com/bulbosaur/game-of-life/config"
	"github.com/bulbosaur/game-of-life/internal/application"
)

func main() {
	cfg, err := config.GettingConfig("..\\..\\config\\config.json")
	if err != nil {
		log.Fatal("Config error:", err)
	}

	app := application.New(cfg)

	ctx := context.Background()
	exitCode := app.Run(ctx)

	os.Exit(exitCode)
}
