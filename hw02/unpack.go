package hw02

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	sitem := []rune(s)
	var resulting_string string
	var n int

	for i, item := range sitem {
		if unicode.IsDigit(item) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(sitem[i-1]) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) {
			n = int(item - '0')
			if n == 0 {
				resulting_string = resulting_string[:len(resulting_string)-1]
				continue
			}
			for j := 0; j < n-1; j++ {
				resulting_string += string(sitem[i-1])
			}
			continue
		}
		resulting_string += string(item)
	}

	return resulting_string, nil
}
