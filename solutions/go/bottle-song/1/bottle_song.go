package bottlesong

import "strings"

func Recite(startBottles, takeDown int) []string {
	var result []string

	for i := 0; i < takeDown; i++ {
		n := startBottles - i
		next := n - 1
		if next < 0 {
			next = 0
		}

		line1 := capitalize(numberWord(n)) + " green " + bottleWord(n) + " hanging on the wall,"
		line2 := capitalize(numberWord(n)) + " green " + bottleWord(n) + " hanging on the wall,"
		line3 := "And if one green bottle should accidentally fall,"
		line4 := "There'll be " + numberWord(next) + " green " + bottleWord(next) + " hanging on the wall."

		result = append(result, line1, line2, line3, line4)

		if i < takeDown-1 {
			result = append(result, "")
		}
	}

	return result
}

func bottleWord(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func numberWord(n int) string {
	words := []string{
		"no", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine", "ten",
	}
	if n < 0 {
		n = 0
	}
	if n > 10 {
		return "many"
	}
	return words[n]
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
