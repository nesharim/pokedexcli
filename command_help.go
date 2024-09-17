package main

import "fmt"

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")

	availableCommands := getCommands()

	for key, value := range availableCommands {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	fmt.Println()

	return nil
}
