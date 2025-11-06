package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	seen := make(map[int]bool)

	for _, d := range divisors {
		if d == 0 {
			continue // avoid division by zero
		}
		for multiple := d; multiple < limit; multiple += d {
			if !seen[multiple] {
				sum += multiple
				seen[multiple] = true
			}
		}
	}

	return sum
}
