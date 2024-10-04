package main

import (
	"errors"
	"fmt"
)

func commandInspectPokemon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Needs a name to inspect a pokemon")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not yet caught this pokemon")
	}

	fmt.Printf("Name: %s", pokemon.Name)
	fmt.Printf("Height: %d", pokemon.Height)
	fmt.Printf("Weight: %d", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}
