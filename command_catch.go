package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatchPokemon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Needs a name to catch a pokemon")
	}

	name := args[0]

	pokemon, err := cfg.pokemonClient.Pokemon(name)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s\n", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon

	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}
