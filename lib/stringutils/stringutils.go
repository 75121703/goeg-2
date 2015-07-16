package stringutils

import (
	"bytes"
	"strings"
	"unicode"
)

// SimpleSimplifyWhitespace accpets an string and normalizes whitespace, that is, to get rid of any
// leading and trailing whitespace and replace each internal sequence of one or more whitespace characters
// with a single space. It then returns the normalized string.
func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

// SimplifyWhitespace accpets an string and normalizes whitespace, that is, to get rid of any leading and
// trailing whitespace and replace each internal sequence of one or more whitespace characters with a single
// space. It then returns the normalized string. More efficient implementation than SimpleSimplifyWhitespace.
func SimplifyWhitespace(s string) string {
	var buffer bytes.Buffer
	skip := true

	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				_, _ = buffer.WriteRune(' ')
				skip = true
			}
		} else {
			_, _ = buffer.WriteRune(char)
			skip = false
		}
	}

	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}

	return s
}

// IsHexDigit checks if char is a hex digit
func IsHexDigit(char rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, char)
}
