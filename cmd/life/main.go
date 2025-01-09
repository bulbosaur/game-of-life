package main

import (
	"fmt"
	"time"

	"github.com/bulbosaur/game-of-life/pkg/life"
)

func main() {
	height := 10
	width := 10
	currentWorld := life.NewWorld(height, width)
	nextWorld := life.NewWorld(height, width)
	currentWorld.Seed()
	for {
		fmt.Println(currentWorld.String())
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
	}
}
