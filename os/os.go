package os

import (
	"errors"
	"os"
	"strings"
)

// IsFileSystemCaseSensitive determines if the file
// system is case sensitive or not.
// It creates a temporary file to check if the file
// can be found in a case-insensitive manner.
// Returned error may be from creating the temporary file,
// closing the file, a call to os.Stat, or os.Remove.
func IsFileSystemCaseSensitive() (isCaseSenstive bool, err error) {
	file, err := os.CreateTemp(os.TempDir(), "gonzal-golib-os-IsFileSystemCaseSensitive-*.TMP")
	if err != nil {
		return false, err
	}
	err = file.Close()
	defer func() {
		if err == nil {
			err = os.Remove(file.Name())
		}
	}()

	if _, err := os.Stat(strings.ToLower(file.Name())); errors.Is(err, os.ErrNotExist) {
		isCaseSenstive = true
	}
	return isCaseSenstive, err
}
