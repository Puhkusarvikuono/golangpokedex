package main
import (
	"fmt"
	"bufio"
	"os"
	"github.com/Puhkusarvikuono/golangpokedex/cleaner"
) 

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		text := scanner.Text()
		if text == "" { 
			fmt.Println("Please input a command.")
			continue
		}
		cleanText := cleaner.CleanInput(text)
		fmt.Printf("Your command was: %s\n", cleanText[0])
	}
}
