package main
import (
	"github.com/Puhkusarvikuono/golangpokedex/internal/pokeapi"
	"time"	
)

type Config struct {
	Next			*string
	Previous	*string
	Pokeapi 	*pokeapi.Client
}

func main() {
	cacheInterval := 5 * time.Second
	cfg := Config{
		Next: nil,
		Previous: nil,
		Pokeapi: pokeapi.NewClient(cacheInterval),
	}
	startRepl(&cfg)
}
