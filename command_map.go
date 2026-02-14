package main

import (
	"fmt"
	"encoding/json"
)
// We'll pop these in the api files 
type response struct {
	Next					*string			`json:"next"`
	Previous			*string			`json:"previous"`
	Results				[]result		`json:"results"`
}

type result struct {
	Name					string			`json:"name"`
}

func commandMap(c *Config) error {

	dat, err := c.Pokeapi.GetLocationPage(c.Next)
	if err != nil {
		return err
	}

	r := response{}
	err = json.Unmarshal(dat, &r)
	if err != nil {
		return err
	}
	for _, item := range r.Results {
		fmt.Println(item.Name)
	}


	c.Next = r.Next
	c.Previous = r.Previous
	return nil
}

