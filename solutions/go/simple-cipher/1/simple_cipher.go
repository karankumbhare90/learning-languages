package cipher

import (
	"strings"
	"unicode"
)

type shift struct {
	distance int
}

type vigenere struct {
	key []int
}

// Normalize input: lowercase only letters A-Z
func normalize(input string) string {
	var sb strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return sb.String()
}

// shiftRune shifts a rune by distance (positive or negative)
func shiftRune(r rune, distance int) rune {
	offset := int(r-'a') + distance
	offset = (offset%26 + 26) % 26 // ensure wraparound
	return rune('a' + offset)
}

/* ---------------- SHIFT CIPHER ---------------- */

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift{distance}
}

func (c shift) Encode(input string) string {
	input = normalize(input)
	var sb strings.Builder
	for _, r := range input {
		sb.WriteRune(shiftRune(r, c.distance))
	}
	return sb.String()
}

func (c shift) Decode(input string) string {
	input = normalize(input)
	var sb strings.Builder
	for _, r := range input {
		sb.WriteRune(shiftRune(r, -c.distance))
	}
	return sb.String()
}

/* ---------------- VIGENERE CIPHER ---------------- */

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	shifts := make([]int, len(key))
	allA := true
	for i, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		shifts[i] = int(r - 'a')
		if r != 'a' {
			allA = false
		}
	}
	if allA {
		return nil
	}
	return vigenere{shifts}
}

func (v vigenere) Encode(input string) string {
	input = normalize(input)
	var sb strings.Builder
	for i, r := range input {
		shift := v.key[i%len(v.key)]
		sb.WriteRune(shiftRune(r, shift))
	}
	return sb.String()
}

func (v vigenere) Decode(input string) string {
	input = normalize(input)
	var sb strings.Builder
	for i, r := range input {
		shift := v.key[i%len(v.key)]
		sb.WriteRune(shiftRune(r, -shift))
	}
	return sb.String()
}
