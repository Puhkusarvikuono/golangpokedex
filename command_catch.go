package main

import (
	"fmt"
	"strings"
	"errors"
	"math/rand"
)

func commandCatch(c *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please add target pokemon name. See 'help' for command usage.")
	}
	target := strings.ToLower(args[0])

	res, err := c.Pokeapi.FetchPokemonData(target)
	if err != nil {
		message := fmt.Sprintf("Could not find a Pokemon named %s", target)
		return errors.New(message)
	}	

	fmt.Printf("Throwing a Pokeball at %s...\n", target)
	// random chance, but now we master ball	
	catchRate := 50
	if res.BaseExperience != 0 {
		catchRate += res.BaseExperience
	} else {
		return errors.New("Base experience does not exist?")
	}

	catchAttempt := rand.Intn(1000)
	catchSuccess := catchAttempt > catchRate
	if catchSuccess {
	fmt.Printf("%s was caught!\n", target)
	c.Pokedex[target] = res
	} else {
		fmt.Printf("%s escaped!\n", target)
	}

	return nil
}
