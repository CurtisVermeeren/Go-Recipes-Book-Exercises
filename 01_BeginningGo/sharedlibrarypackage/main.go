package sharedlibrarypackage

import (
	"strings"
	"unicode"
)

// ToUpperCase returns the string changed to uppercase
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ToLowerCase returns the string changed to lowercase
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// ToFirstUpper returns the string with the first letter uppercase
func ToFirstUpper(s string) string {
	if len(s) < 1 {
		return s
	}
	s = strings.Trim(s, " ")
	s = strings.ToLower(s)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
