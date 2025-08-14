package luhn

import "unicode"

func Valid(id string) bool {
	var digits []int
	for _, r := range id {
		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		digits = append(digits, int(r-'0'))
	}

	// Length must be > 1 after removing spaces
	if len(digits) <= 1 {
		return false
	}

	// Apply Luhn algorithm
	sum := 0
	double := false
	// Start from rightmost digit
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}

	return sum%10 == 0
}
