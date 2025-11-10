package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	// Normalize
	var sb strings.Builder
	for _, r := range strings.ToLower(pt) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}
	text := sb.String()
	if len(text) == 0 {
		return ""
	}

	n := len(text)
	c := int(math.Ceil(math.Sqrt(float64(n))))
	r := int(math.Ceil(float64(n) / float64(c)))

	// Build rows
	rows := make([]string, r)
	for i := 0; i < r; i++ {
		start := i * c
		end := start + c
		if end > n {
			end = n
		}
		rows[i] = text[start:end]
	}

	// Build columns
	var result []string
	for col := 0; col < c; col++ {
		var colStr strings.Builder
		for row := 0; row < r; row++ {
			if col < len(rows[row]) {
				colStr.WriteByte(rows[row][col])
			} else {
				colStr.WriteByte(' ')
			}
		}
		result = append(result, colStr.String())
	}

	return strings.Join(result, " ")
}
