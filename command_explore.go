package main

import (
	"fmt"
	"errors"
)

func commandExplore(c *Config, args []string) error {

	if len(args) == 0 {
		return errors.New("Please provide argument. Type 'help' for commands.")
	}
	
	targetArea := args[0]
	res, err := c.Pokeapi.ExploreLocationResponse(targetArea)
	if err != nil {
		return err
	}	

	fmt.Printf("Exploring area: %s... \nFound Pokemon:\n", targetArea)

	for _, item := range res.PokemonEncounters {
		fmt.Print("- ")
		fmt.Println(item.Pokemon.Name)
	}

	return nil
}
