package series

// All returns all contiguous substrings of s of length n.
func All(n int, s string) []string {
	if n > len(s) {
		return []string{} // or panic, depending on your design
	}

	result := make([]string, 0, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

// UnsafeFirst returns the first substring of s with length n.
// It’s “unsafe” because it doesn’t check for n > len(s).
func UnsafeFirst(n int, s string) string {
	return s[:n]
}
