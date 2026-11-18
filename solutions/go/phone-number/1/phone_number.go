package phonenumber

import (
	"errors"
	"strings"
	"unicode"
)

func isValidDigit(n byte) bool {
	return n >= '2' && n <= '9'
}

func Number(phoneNumber string) (string, error) {
	var clearNumber strings.Builder

	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			clearNumber.WriteRune(r)
		}
	}

	if clearNumber.Len() < 10 || clearNumber.Len() > 11 {
		return "", errors.New("invalid format, number must be 10 chars")
	} else if clearNumber.Len() == 11 {
		if clearNumber.String()[0] != '1' {
			return "", errors.New("invalid country code, code must be 1")
		}
		if !isValidDigit(clearNumber.String()[1]) || !isValidDigit(clearNumber.String()[4]) {
			return "", errors.New("invalid digit")
		}
		return clearNumber.String()[1:], nil
	}

	if !isValidDigit(clearNumber.String()[0]) || !isValidDigit(clearNumber.String()[3]) {
		return "", errors.New("invalid digit")
	}

	return clearNumber.String(), nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return number, err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return number, err
	}
	var formattedNumber strings.Builder

	for i, r := range number {
		if i == 0 {
			formattedNumber.WriteString("(")
		} else if i == 3 {
			formattedNumber.WriteString(") ")
		} else if i == 6 {
			formattedNumber.WriteString("-")
		}
		formattedNumber.WriteRune(r)
	}
	return formattedNumber.String(), nil
}