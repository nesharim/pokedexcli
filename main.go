package main

import (
	"time"

	"github.com/nesharim/pokedexcli/internal/pokemonapi"
)

type config struct {
	nextLocationAreasURL     *string
	previousLocationAreasURL *string
	caughtPokemon            map[string]pokemonapi.Pokemon
	pokemonClient            pokemonapi.Client
}

func main() {
	cfg := config{
		pokemonClient: pokemonapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokemonapi.Pokemon),
	}

	repl(&cfg)
}
