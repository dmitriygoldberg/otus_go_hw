package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var unpackedString strings.Builder

	if !utf8.ValidString(input) {
		return "", ErrInvalidString
	}

	var prev rune
	for _, r := range input {
		if unicode.IsDigit(r) {
			if !isPrevLetter(prev) {
				return "", ErrInvalidString
			}

			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrInvalidString
			}

			unpackedString.WriteString(strings.Repeat(string(prev), count))
		} else if isPrevLetter(prev) {
			unpackedString.WriteRune(prev)
		}

		prev = r
	}

	if isPrevLetter(prev) {
		unpackedString.WriteRune(prev)
	}

	return unpackedString.String(), nil
}

func isPrevLetter(prev rune) bool {
	return prev > 0 && !unicode.IsDigit(prev)
}
