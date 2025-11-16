package resistorcolortrio

import "fmt"

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Label(colors []string) string {
	// First two digits give base value
	first := colorMap[colors[0]]
	second := colorMap[colors[1]]
	baseValue := first*10 + second

	// Third color specifies the exponent of 10
	exp := colorMap[colors[2]]

	// Simple integer power of 10
	value := baseValue
	for i := 0; i < exp; i++ {
		value *= 10
	}

	// Scale units in steps of 1000
	unit := "ohms"
	display := value

	if display >= 1_000_000_000 {
		display /= 1_000_000_000
		unit = "gigaohms"
	} else if display >= 1_000_000 {
		display /= 1_000_000
		unit = "megaohms"
	} else if display >= 1_000 {
		display /= 1_000
		unit = "kiloohms"
	}

	return fmt.Sprintf("%d %s", display, unit)
}
