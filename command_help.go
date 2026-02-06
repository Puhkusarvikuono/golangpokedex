package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}
