// Package proverb implements the generation of the "For want of a nail..." proverb.
package proverb

// Proverb generates a proverbial rhyme from the given list of strings.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	var result []string
	// Generate lines for pairs
	for i := 0; i < len(rhyme)-1; i++ {
		line := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		result = append(result, line)
	}

	// Add final line
	result = append(result, "And all for the want of a "+rhyme[0]+".")

	return result
}
