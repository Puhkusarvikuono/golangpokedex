package main

import (
	"fmt"
	"errors"
	"strings"
)

func commandInspect(c *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please add target pokemon name. See 'help' for command usage.")
	}
	target := strings.ToLower(args[0])

	dexPokemon, exists := c.Pokedex[target]
	fmt.Println(dexPokemon)
	if exists {
		fmt.Printf("Name: %s\n", dexPokemon.Name)
		fmt.Printf("Weight: %d\n", dexPokemon.Weight)
		fmt.Printf("Height: %d\n", dexPokemon.Height)
		fmt.Printf("Stats:\n")	
		for _, stat := range dexPokemon.Stats {
			fmt.Printf("		-%s: %d\n", stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, stat := range dexPokemon.Types.Type {
			fmt.Printf("		-%s: %d\n", stat.Name)
		}
	} else {
		return errors.New("you have not caught that pokemon")
	}

	return nil

}


