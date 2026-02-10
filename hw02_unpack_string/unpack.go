package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var outputString strings.Builder

	runes := []rune(inputString)

	if len(runes) == 0 {
		return outputString.String(), nil
	}

	for i := 0; i < len(runes); i++ {
		if isDigit(runes[i]) {
			if i == 0 || isDigit(runes[i-1]) {
				return "", ErrInvalidString
			}
			for j := 0; j < int(runes[i]-'0'); j++ {
				outputString.WriteRune(runes[i-1])
			}
		} else {
			if i+1 >= len(runes) {
				outputString.WriteRune(runes[i])
			} else if !isDigit(runes[i+1]) {
				outputString.WriteRune(runes[i])
			}
		}
	}

	return outputString.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
