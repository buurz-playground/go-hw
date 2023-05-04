package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev rune
	var noNumberEscape bool
	var escape bool

	for i, r := range s {
		if escape {
			if r != '\\' && !unicode.IsDigit(r) {
				return "", ErrInvalidString
			}

			result.WriteRune(r)
			noNumberEscape = false
			prev = r
			escape = false
			continue
		}

		if r == '\\' {
			noNumberEscape = false
			escape = true
			prev = r
			continue
		}

		switch {
		case unicode.IsDigit(r):
			if i == 0 || unicode.IsDigit(prev) && noNumberEscape {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(r))
			if count == 0 {
				str := result.String()
				if len(str) > 0 {
					str = str[:len(str)-1]
					result.Reset()
					result.WriteString(str)
				}
				continue
			}
			result.WriteString(strings.Repeat(string(prev), count-1))
			noNumberEscape = true
		case unicode.IsLetter(r):
			noNumberEscape = false
			result.WriteRune(r)

		default:
			return "", ErrInvalidString
		}

		prev = r
	}

	if escape {
		return "", ErrInvalidString
	}

	return result.String(), nil
}
