package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bulbosaur/game-of-life/config"
	"github.com/bulbosaur/game-of-life/pkg/life"
)

func main() {
	var (
		currentWorld, nextWorld *life.World
	)

	cfg, err := config.GettingConfig("..\\..\\config\\config.json")
	if err != nil {
		log.Fatal("Config error:", err)
	}

	// if !life.EmptyStateFile(cfg.StatePath) {
	// 	err = currentWorld.LoadState(cfg.StatePath)
	// 	if err != nil {

	// 	}
	// } else {
	currentWorld = life.NewWorld(cfg.Height, cfg.Width)
	nextWorld = life.NewWorld(cfg.Height, cfg.Width)
	currentWorld.Seed()
	// }
	for {
		currentWorld.SaveState(cfg.StatePath)
		fmt.Println(currentWorld.String())
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
	}
}
