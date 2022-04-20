package stderr

import (
	"fmt"
	"os"
)

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stderr, a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}
