package main

import (
	"fmt"
	"errors"
)

func commandExplore(c *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide argument. Type 'help' for commands.")
	}
	fmt.Println("Exploring: ", args[0])
	return nil
}

