package main

import (
	"fmt"
	"errors"
)

func commandPokedex(c *Config, args []string) error {
	if len(c.Pokedex) == 0 {
		return errors.New("Pokedex is empty. Go catch a pokemon!")
	}
	fmt.Println("Your Pokedex:")
	for key := range c.Pokedex {
		fmt.Printf("		- %s\n", key)
	}
	return nil	
}
