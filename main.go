package main

import (
	"time"

	"github.com/nesharim/pokedexcli/internal/pokemonapi"
)

type config struct {
	nextLocationAreasURL     *string
	previousLocationAreasURL *string
	pokemonClient            pokemonapi.Client
}

func main() {
	cfg := config{
		pokemonClient: pokemonapi.NewClient(time.Hour),
	}

	repl(&cfg)
}
