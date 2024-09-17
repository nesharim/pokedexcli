package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokemonClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}

	for _, location := range resp.Results {
		fmt.Printf("%s-area\n", location.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.previousLocationAreasURL = resp.Previous

	return nil
}
