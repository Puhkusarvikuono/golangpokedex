package main

import (
	"fmt"
)

func commandHelp(c *Config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, item := range commands {
		fmt.Printf("Command name: %s \nDescription: %s \n", item.name, item.description)
	}
	return nil
}
