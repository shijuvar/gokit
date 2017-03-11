// Package strutils provides string utility functions
package strutils

import (
	"strings"
	"unicode"
)

// Returns the string changed with upper case.
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// Returns the string changed with lower case.
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// Returns the string changed to upper case for its first letter.
func ToFirstUpper(s string) string {
	if len(s) < 1 { // if the empty string
		return s
	}
	// Trim the string
	t := strings.Trim(s, " ")
	// Convert all letters to lower case
	t = strings.ToLower(t)
	res := []rune(t)
	// Convert first letter to upper case
	res[0] = unicode.ToUpper(res[0])
	return string(res)
}
