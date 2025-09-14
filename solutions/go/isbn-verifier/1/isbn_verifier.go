package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	digit := []rune{}

    for _, r := range isbn {
        if r == '-' {
            continue
        }

        digit = append(digit, r)
    }

    if len(digit) != 10{
        return false
    }

    sum := 0

    for i, r := range digit{
        position := 10 - i
        var value int

        if i == 9 && (r == 'X' || r == 'x'){
            value = 10
        } else if unicode.IsDigit(r){
            value = int(r - '0')
        } else{
            return false
        }

        sum += value * position
    }

    return sum % 11 == 0
}
