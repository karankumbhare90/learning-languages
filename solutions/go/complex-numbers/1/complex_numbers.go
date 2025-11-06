package complexnumbers

import "math"

// Number represents a complex number in the form a + b*i.
type Number struct {
	real, imag float64
}

// Real returns the real part of the complex number.
func (n Number) Real() float64 {
	return n.real
}

// Imaginary returns the imaginary part of the complex number.
func (n Number) Imaginary() float64 {
	return n.imag
}

// Add returns the sum of two complex numbers.
func (n1 Number) Add(n2 Number) Number {
	return Number{
		real: n1.real + n2.real,
		imag: n1.imag + n2.imag,
	}
}

// Subtract returns the difference between two complex numbers.
func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real: n1.real - n2.real,
		imag: n1.imag - n2.imag,
	}
}

// Multiply returns the product of two complex numbers.
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real: n1.real*n2.real - n1.imag*n2.imag,
		imag: n1.imag*n2.real + n1.real*n2.imag,
	}
}

// Times scales a complex number by a real factor.
func (n Number) Times(factor float64) Number {
	return Number{
		real: n.real * factor,
		imag: n.imag * factor,
	}
}

// Divide returns the result of dividing n1 by n2.
func (n1 Number) Divide(n2 Number) Number {
	denom := n2.real*n2.real + n2.imag*n2.imag
	return Number{
		real: (n1.real*n2.real + n1.imag*n2.imag) / denom,
		imag: (n1.imag*n2.real - n1.real*n2.imag) / denom,
	}
}

// Conjugate returns the complex conjugate of n.
func (n Number) Conjugate() Number {
	return Number{
		real: n.real,
		imag: -n.imag,
	}
}

// Abs returns the absolute value (magnitude) of n.
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imag*n.imag)
}

// Exp returns e raised to the power of the complex number n.
func (n Number) Exp() Number {
	expReal := math.Exp(n.real)
	return Number{
		real: expReal * math.Cos(n.imag),
		imag: expReal * math.Sin(n.imag),
	}
}
