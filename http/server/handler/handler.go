package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bulbosaur/game-of-life/config"
	"github.com/bulbosaur/game-of-life/internal/service"
)

type Decorator func(http.Handler) http.Handler

type LifeStates struct {
	service.LifeService
}

func New(ctx context.Context,
	lifeService service.LifeService,
) (http.Handler, error) {
	serveMux := http.NewServeMux()

	lifeState := LifeStates{
		LifeService: lifeService,
	}

	serveMux.HandleFunc("/nextstate", lifeState.nextState)

	return serveMux, nil
}

func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- {
		decorated = ds[d](decorated)
	}

	return decorated
}

func (ls *LifeStates) nextState(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.GettingConfig("..\\..\\config\\config.json")
	if err != nil {
		log.Fatal("Config error:", err)
	}

	worldState := ls.LifeService.NewState()

	err = worldState.SaveState(cfg.StatePath)
	if err != nil {
		log.Fatal("Saving is faled")
	}

	worldString := worldState.String()

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, worldString)
}
