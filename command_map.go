package main

import (
	"fmt"
	"os"
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
	// Temporary dummy file to mimic the real API
	dat, err := os.ReadFile("./cache.json")
	if err != nil {
		return err
	}
	// API call here, we get data, pass on the config
	/* decide  which URL to use (cfg.Next / cfg.Previous / nil)
call cfg.PokeClient.ListLocationAreas(url)
get some data 
print names from data 
store cfg.Next/cfg.Previous from the response */

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

