package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
        return "", errors.New("input must be between 1 to 3999")
    }

    vals := []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}

    result := ""

    for i := 0; i < len(vals); i++{
        for input >= vals[i]{
            result += symbols[i]
            input -= vals[i]
        }
    }

    return result, nil
}
