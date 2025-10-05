package armstrong

// IsNumber returns true if n is an Armstrong (narcissistic) number.
// Example: 153 -> 1^3 + 5^3 + 3^3 == 153  => true
func IsNumber(n int) bool {
	if n < 0 {
		return false
	}
	// Single-digit numbers 0..9 are Armstrong numbers by definition
	if n >= 0 && n <= 9 {
		return true
	}

	// Count digits and collect digits
	num := n
	digits := make([]int, 0, 10)
	for num > 0 {
		d := num % 10
		digits = append(digits, d)
		num /= 10
	}
	k := len(digits)

	// integer power function (digit^k)
	pow := func(base, exp int) int {
		res := 1
		for i := 0; i < exp; i++ {
			res *= base
		}
		return res
	}

	sum := 0
	for _, d := range digits {
		sum += pow(d, k)
		// early exit if sum already exceeds n (optional optimization)
		if sum > n {
			return false
		}
	}

	return sum == n
}
