package phonenumber

import (
	"errors"
	"regexp"
	"strings"
)

// Number cleans and validates the phone number
func Number(phoneNumber string) (string, error) {
	// Remove all non-digit characters
	re := regexp.MustCompile(`\d`)
	digits := strings.Join(re.FindAllString(phoneNumber, -1), "")

	// Remove country code if present
	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	}

	// Validate length
	if len(digits) != 10 {
		return "", errors.New("invalid number length")
	}

	// Validate NANP rules
	if digits[0] < '2' || digits[0] > '9' {
		return "", errors.New("invalid area code")
	}
	if digits[3] < '2' || digits[3] > '9' {
		return "", errors.New("invalid exchange code")
	}

	return digits, nil
}

// AreaCode returns the 3-digit area code
func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format returns a standardized phone number format
func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + num[:3] + ") " + num[3:6] + "-" + num[6:], nil
}
