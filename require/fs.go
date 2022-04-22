package require

import (
	"fmt"

	"github.com/algonzalez/golib/fs"
)

// DirExists enforces that the dir path already exists.
func DirExists(path string) {
	if !fs.DirExists(path) {
		panic(fmt.Sprintf("could not find directory '%s'", path))
	}
}

// FileExists enforces that the file path already exists.
func FileExists(path string) {
	if !fs.FileExists(path) {
		panic(fmt.Sprintf("could not find file '%s'", path))
	}
}
