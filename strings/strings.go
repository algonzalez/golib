// Package strings provides utility functions for working with strings.
package strings

import (
	"fmt"
	"strconv"
	"strings"
)

// PadCenter returns a string with center-aligned text that has been
// padded with spaces on both ends to reach the specified length.
func PadCenter(s string, length int) string {
	if len(s) >= length {
		return s
	}
	padWidth := (length - len(s)) / 2
	s = PadLeft(s, len(s)+padWidth)
	return PadRight(s, length)
}

// PadCenterWith returns a string with center-aligned text that has been
// padded with specified padChar on both ends to reach the specified length.
func PadCenterWith(s string, length int, padChar byte) string {
	if len(s) >= length {
		return s
	}
	padWidth := (length - len(s)) / 2
	s = PadLeftWith(s, len(s)+padWidth, padChar)
	return PadRightWith(s, length, padChar)
}

// PadLeft returns a string with right-aligned text that has been
// padded with spaces to reach the specified length.
func PadLeft(s string, length int) string {
	if len(s) >= length {
		return s
	}
	padformat := "%" + strconv.Itoa(length) + "s"
	return fmt.Sprintf(padformat, s)
}

// PadLeftWith returns a string with right-aligned text that has been
// padded with specified padChar to reach the specified length.
func PadLeftWith(s string, length int, padChar byte) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(padChar), length-len(s))
	return fmt.Sprintf("%v%v", padding, s)
}

// PadRight returns a string with left-aligned text that has been
// padded with spaces to reach the specified length.
func PadRight(s string, length int) string {
	if len(s) >= length {
		return s
	}
	padformat := "%-" + strconv.Itoa(length) + "s"
	return fmt.Sprintf(padformat, s)
}

// PadRightWith returns a string with left-aligned text that has been
// padded with specified padChar to reach the specified length.
func PadRightWith(s string, length int, padChar byte) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(padChar), length-len(s))
	return fmt.Sprintf("%v%v", s, padding)
}
