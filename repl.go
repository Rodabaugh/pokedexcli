package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanUserInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}

		command.callback()
	}
}

func cleanUserInput(input string) []string {
	lowerInput := strings.ToLower(input)
	inputWords := strings.Fields(lowerInput)
	return inputWords
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get a map",
			callback:    commandMap,
		},
	}
}
