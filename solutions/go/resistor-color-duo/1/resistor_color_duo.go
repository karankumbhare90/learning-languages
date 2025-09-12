package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	colorMap := map[string]int{
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

	if len(colors) < 2 {
		return 0 // not enough colors to form a value
	}

	first := colorMap[colors[0]]
	second := colorMap[colors[1]]

	return first*10 + second
}
