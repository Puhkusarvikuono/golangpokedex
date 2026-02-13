package main

type Config struct {
	Next			*string
	Previous	*string
}

func main() {
	cfg := Config{
		Next: nil,
		Previous: nil,
		//possible client
	}
	startRepl(&cfg)
}
