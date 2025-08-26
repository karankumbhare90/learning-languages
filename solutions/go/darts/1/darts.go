package darts

import "math"

// Score returns the points earned for a dart landing at (x, y).
func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)

	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}