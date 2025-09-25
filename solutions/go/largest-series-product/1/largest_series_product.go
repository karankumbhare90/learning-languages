package lsproduct
import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return 0, errors.New("span must not be bigger than string length")
	}
	if span == 0 {
		return 1, nil
	}

    // Check digits string only contains numbers
	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return 0, errors.New("input must only contain digits")
		}
	}

	var maxProduct int64 = 0

	// Sliding window of size span
	for i := 0; i <= len(digits)-span; i++ {
		product := int64(1)
		for j := 0; j < span; j++ {
			d := digits[i+j] - '0' // convert rune to int
			product *= int64(d)
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
