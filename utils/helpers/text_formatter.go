package helpers

import (
	"strings"
	"unicode"
)

// Capitalize first letter only
func CFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Capitalize first letter of each words
func CWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			// Convert the first rune to uppercase
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			// Convert the rest of the runes to lowercase
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}
