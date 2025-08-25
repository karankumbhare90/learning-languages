package pangram

func IsPangram(input string) bool {
	const all = (1 << 26) - 1
	mask := 0

	for _, r := range input {
		// Normalize to lowercase ASCII if it's A–Z
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		// Only count a–z letters
		if r >= 'a' && r <= 'z' {
			mask |= 1 << (r - 'a')
			if mask == all {
				return true
			}
		}
	}
	return mask == all
}
