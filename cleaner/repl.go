package cleaner
import ("strings")

func CleanInput(text string) []string {
	// cleanInput splits text into lowercase words, trimming whitespace.
	// Note: this could also be implemented with strings.Fields after
	// trimming + lowercasing, but this version is expanded for clarity.

	splitText := strings.Split(text, " ")
	words := []string{}
	for _, s := range splitText {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		if s != "" {
		words = append(words, s)
		}
	}
	return words
}
