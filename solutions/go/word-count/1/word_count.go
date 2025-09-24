package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)

	// Lowercase everything for case insensitivity
	phrase = strings.ToLower(phrase)

	// Regex: word = letters/digits, may contain one or more apostrophe parts
	re := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)*`)

	words := re.FindAllString(phrase, -1)
	for _, w := range words {
		freq[w]++
	}

	return freq
}
