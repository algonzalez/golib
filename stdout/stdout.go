package stdout

import (
	"fmt"
	"os"
)

// Print formats using the default formats for its operands and writes to os.Stdout.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, a...)
}

// Printf formats according to a format specifier and writes to os.Stdout.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

// Println formats using the default formats for its operands and writes to os.Stdout.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}
