package wordy

import (
	"strconv"
	"strings"
)

// Answer parses and evaluates simple math word problems.
func Answer(question string) (int, bool) {
	// Trim question marks and extra spaces
	q := strings.TrimSpace(question)
	q = strings.TrimSuffix(q, "?")

	// Must start with "What is "
	if !strings.HasPrefix(q, "What is ") {
		return 0, false
	}

	// Remove the starting phrase
	q = strings.TrimPrefix(q, "What is ")

	// Replace verbal operators with symbols for clarity
	q = strings.ReplaceAll(q, "multiplied by", "mul")
	q = strings.ReplaceAll(q, "divided by", "div")
	q = strings.ReplaceAll(q, "plus", "add")
	q = strings.ReplaceAll(q, "minus", "sub")

	// Split into words
	tokens := strings.Fields(q)
	if len(tokens) == 0 {
		return 0, false
	}

	// Try parsing the first number
	current, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	i := 1
	for i < len(tokens) {
		if i+1 >= len(tokens) {
			return 0, false // incomplete operation
		}

		op := tokens[i]
		numStr := tokens[i+1]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, false
		}

		switch op {
		case "add":
			current += num
		case "sub":
			current -= num
		case "mul":
			current *= num
		case "div":
			current /= num
		default:
			return 0, false // unsupported operation
		}
		i += 2
	}

	return current, true
}
