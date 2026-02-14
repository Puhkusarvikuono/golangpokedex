package main
import "github.com/Puhkusarvikuono/golangpokedex/internal/pokeapi"

type Config struct {
	Next			*string
	Previous	*string
	Pokeapi 	*pokeapi.Client
}

func main() {
	cfg := Config{
		Next: nil,
		Previous: nil,
		Pokeapi: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
