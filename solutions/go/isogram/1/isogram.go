package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

    for _, r := range word {
		if r == ' ' || r == '-' {
			continue // skip spaces and hyphens
		}
        
        r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
    }
    return true  
}
