// Package triangle provides utilities for classifying triangles by side lengths.
package triangle

import "math"

// Kind represents the classification of a triangle.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the type of triangle given side lengths a, b, c.
//
// Rules:
// - If any side <= 0, or is NaN/Inf → NaT
// - Must satisfy triangle inequality, otherwise NaT
// - All sides equal → Equ
// - At least two sides equal → Iso
// - All sides different → Sca
func KindFromSides(a, b, c float64) Kind {
	// Reject invalid values
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	// Triangle inequality
	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	// Classification
	switch {
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
