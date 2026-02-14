package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	// cleanInput splits text into lowercase words, trimming whitespace.
	// Note: this could also be implemented with strings.Fields after
	// trimming + lowercasing, but this version is expanded for clarity.

	splitText := strings.Split(text, " ")
	words := []string{}
	for _, s := range splitText {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		if s != "" {
		words = append(words, s)
		}
	}
	return words
}

type cliCommand struct {
	name				string
	description	string
	callback		func(*Config) error
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		if rawInput == "" {
			fmt.Println("Please input a command. Type 'help' for help message.")
			continue
		}
		input := cleanInput(rawInput)
		if commandFunc, exists := commands[input[0]]; exists {
			err := commandFunc.callback(cfg)
			if err != nil {
				fmt.Println("Command failed:", err)
			}

		} else {
			fmt.Println("Please input a command. Unknown command:", input[0])
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
}

func getCommands() (map[string]cliCommand) {
	commands := map[string]cliCommand{
		"exit": {
			name: 				"exit",
			description:	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name:					"help",
			description:	"Displays a help message",
			callback:			commandHelp,	
		},
		"map": {
			name:					"map",
			description:	"Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays 20 additional locations.",
			callback:			commandMap,
		},
		"mapb": {
			name: 				"map back",
			description:	"Displays the previous 20 location areas of the map command.",
			callback:			commandMapb,
		},
	}
	return commands
}
