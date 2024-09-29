package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
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
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous page of location areas",
			callback:    commandMapb,
		},
	}
}
