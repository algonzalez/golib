package stdout

import (
	"fmt"
	"os"
)

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}
