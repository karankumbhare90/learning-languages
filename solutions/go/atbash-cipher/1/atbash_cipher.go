package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	// Normalize input: lowercase, strip punctuation
	var cleaned []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			// Atbash transform
			cleaned = append(cleaned, 'a'+'z'-r)
		} else if unicode.IsDigit(r) {
			// keep digits
			cleaned = append(cleaned, r)
		}
	}

	// Group into blocks of 5
	var result strings.Builder
	for i, r := range cleaned {
		if i > 0 && i%5 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}
	return result.String()
}
