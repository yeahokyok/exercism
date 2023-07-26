package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	normalizedWord := strings.ToLower(word)
	for _, char := range normalizedWord {
		if !unicode.IsLetter(char) {
			continue
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
