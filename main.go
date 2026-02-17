package main
import (
	"github.com/Puhkusarvikuono/golangpokedex/internal/pokeapi"
	"time"	
)

type Config struct {
	Next			*string
	Previous	*string
	Pokeapi 	*pokeapi.Client
	Pokedex		map[string]pokeapi.Pokemon
}

func main() {
	cacheInterval := 5 * time.Second
	cfg := Config{
		Next: nil,
		Previous: nil,
		Pokeapi: pokeapi.NewClient(cacheInterval),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
