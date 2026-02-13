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

type Config struct {
	next			string
	previous	string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		next: "",
		previous: "",
	}
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
	}
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
			commandFunc.callback(cfg)
		} else {
			fmt.Println("Please input a command. Unknown command:", input[0])
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
}


