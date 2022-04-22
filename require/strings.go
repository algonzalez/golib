package require

import "strings"

// StringIsEmpty enforces that a string has len of zero.
func StringIsEmpty(s string) {
	if s != "" {
		panic("string must be empty")
	}
}

// StringIsNotEmpty enforces that a string has a len greater than zero.
func StringIsNotEmpty(s string) {
	if s == "" {
		panic("string must not be empty")
	}
}

// StringIsNotWhitespaceOnly enforces that a string is not empty
// or composed of only whitespace.
func StringIsNotWhitespaceOnly(s string) {
	if strings.TrimSpace(s) == "" {
		panic("string must not be empty or composed of only whitespace")
	}
}
