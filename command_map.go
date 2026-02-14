package main

import (
	"fmt"
	"errors"
)

func commandMap(c *Config) error {

	res, err := c.Pokeapi.FetchLocationResponse(c.Next)
	if err != nil {
		return err
	}
	
	c.Next = res.Next
	c.Previous = res.Previous

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}

	return nil
}

func commandMapb(c *Config) error {

	if c.Previous == nil {
		return errors.New("You're on the first page")
	}

	res, err := c.Pokeapi.FetchLocationResponse(c.Previous) 
	if err != nil {
		return err
	}
	
	c.Next = res.Next
	c.Previous = res.Previous

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}

