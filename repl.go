package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	callback    func() error
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
		"exit": {
			callback:    commandExit,
			name:        "exit",
			description: "Exit the Pokedex",
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func repl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		availableCommands := getCommands()

		if len(cleaned) == 0 {
			fmt.Println("Please enter a command")
			commandHelp()
			continue
		}

		command, ok := availableCommands[text]

		if !ok {
			fmt.Print("Invalid Command\n\n")
			continue
		}

		command.callback()
		fmt.Println()

	}
}
