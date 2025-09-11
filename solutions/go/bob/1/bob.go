package bob

import (
	"strings"
	"unicode"
)

// Hey returns Bob's response to a remark.
func Hey(remark string) string {
	// Trim surrounding whitespace
	trimmed := strings.TrimSpace(remark)

	// Silence case
	if trimmed == "" {
		return "Fine. Be that way!"
	}

	// Check if it's a question
	isQuestion := strings.HasSuffix(trimmed, "?")

	// Check if it's yelling:
	// At least one letter must be present and all letters must be uppercase.
	hasLetter := false
	isAllCaps := true
	for _, r := range trimmed {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				isAllCaps = false
			}
		}
	}
	isYelling := hasLetter && isAllCaps

	// Decision tree
	switch {
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
