package main

import (
	"bufio"
	"fmt"
	"os"
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

func commandHelp() error {
	fmt.Print("\n")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")

	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func cliCommand(cmd string) {
	command, ok := getCommands()[cmd]
	if !ok {
		panic("cmd does not exist")
	}

	command.callback()
}

func prompt() {
	for {

		fmt.Print("Pokedex > ")

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			key := scanner.Text()
			cliCommand(key)
			fmt.Println()
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading i input: \n", err)
			break
		}

	}
}

func main() {
	prompt()
}
