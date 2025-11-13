package allyourbase

import "errors"

// ConvertToBase converts a number from one base to another.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// --- Validation ---
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// Remove leading zeros
	for len(inputDigits) > 0 && inputDigits[0] == 0 {
		inputDigits = inputDigits[1:]
	}

	// Empty input = zero
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	// Check all digits are valid
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	// --- Step 1: Convert input to base-10 integer ---
	value := 0
	for _, d := range inputDigits {
		value = value*inputBase + d
	}

	// --- Step 2: Convert base-10 integer to outputBase ---
	if value == 0 {
		return []int{0}, nil
	}

	var output []int
	for value > 0 {
		remainder := value % outputBase
		output = append([]int{remainder}, output...) // prepend
		value /= outputBase
	}

	return output, nil
}
