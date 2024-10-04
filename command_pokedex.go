package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")

	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
