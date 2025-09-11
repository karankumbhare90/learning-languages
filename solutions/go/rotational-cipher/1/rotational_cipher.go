package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	shift := rune(shiftKey%26)
    results := []rune{}

    for _,r := range plain{
        if unicode.IsUpper(r){
            newRune := 'A' + (r - 'A' + shift) % 26	
            results = append(results, newRune)
        } else if unicode.IsLower(r){
            newRune := 'a' + (r - 'a' + shift) % 26	
            results = append(results, newRune)
        } else {
            results = append(results, r)
        }
    }

    return string(results)
}
