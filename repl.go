package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	callback    func(cfg *config, args ...string) error
	name        string
	description string
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			callback:    commandHelp,
			name:        "help",
			description: "Displays a help message",
		},
		"explore": {
			callback:    commandExplore,
			name:        "explore",
			description: "Explores an area of the map",
		},
		"map": {
			callback:    commandMap,
			name:        "map",
			description: "List next 20 locations",
		},
		"mapb": {
			callback:    commandMapB,
			name:        "mapb",
			description: "List previous 20 locations",
		},
		"exit": {
			callback:    commandExit,
			name:        "exit",
			description: "Exit the Pokedex",
		},
		"catch": {
			callback:    commandCatchPokemon,
			name:        "catch <pokemon-name>",
			description: "Gives you an opportunity to catch a pokemon",
		},
		"inspect": {
			callback:    commandInspectPokemon,
			name:        "inspect <pokemon-name>",
			description: "Allows you to inspect a pokemon you have caught",
		},
		"pokedex": {
			callback:    commandPokedex,
			name:        "pokedex",
			description: "View all pokemon in the pokedex",
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func repl(cfg *config) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		args := []string{}
		cleaned := cleanInput(text)

		availableCommands := getCommands()

		if len(cleaned) == 0 {
			fmt.Println("Please enter a command")
			commandHelp(cfg)
			continue
		}

		option := cleaned[0]

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := availableCommands[option]

		if !ok {
			fmt.Print("Invalid Command\n\n")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println()

	}
}
