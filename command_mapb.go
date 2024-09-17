package main

import (
	"fmt"
)

func commandMapB(cfg *config) error {
	if cfg.previousLocationAreasURL == nil {
		return fmt.Errorf("Previous URL null")
	}
	resp, err := cfg.pokemonClient.ListLocationAreas(cfg.previousLocationAreasURL)
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
