package anagram

import (
	"sort"
	"strings"
)

// helper: sort letters in a string
func sortedString(s string) string {
	// convert to slice of runes for UTF-8 support
	runes := []rune(strings.ToLower(s))
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	var result []string
	subjectLower := strings.ToLower(subject)
	subjectSorted := sortedString(subject)

	for _, cand := range candidates {
		candLower := strings.ToLower(cand)

		// skip if candidate is same as subject
		if candLower == subjectLower {
			continue
		}

		// check if sorted letters match
		if sortedString(cand) == subjectSorted {
			result = append(result, cand)
		}
	}
	return result
}
