package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a phrase into its acronym.
// Example: "Portable Network Graphics" -> "PNG"
func Abbreviate(s string) string {
	// Replace hyphens with spaces so they’re treated as separators
	s = strings.ReplaceAll(s, "-", " ")

	var acronym strings.Builder
	words := strings.Fields(s)

	for _, word := range words {
		for _, r := range word {
			if unicode.IsLetter(r) {
				acronym.WriteRune(unicode.ToUpper(r))
				break // move to next word after first letter
			}
		}
	}

	return acronym.String()
}
