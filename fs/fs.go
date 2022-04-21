package fs

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

// DirExists returns true if the directory path was found
// and is in fact a directory; otherwise, it returns false
func DirExists(path string) bool {
	info, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist) && info.IsDir()
}

// FileExists returns true if the file path was found
// and is in fact a file; otherwise, it returns false
func FileExists(path string) bool {
	info, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist) && !info.IsDir()
}

// GetFileInfo retuns an fs.FileInfo object for the file path; if found.
func GetFileInfo(path string) (info fs.FileInfo, exists bool) {
	info, err := os.Stat(path)
	return info, !errors.Is(err, fs.ErrNotExist) && !info.IsDir()
}

// IsCaseSensitive determines if the file system is case sensitive or not.
func IsCaseSensitive() (bool, error) {
	file, err := os.CreateTemp(os.TempDir(), "gonzal-golib-fs-IsCaseSensitive-*.TMP")
	if err != nil {
		return false, err
	}
	closeError := file.Close()
	defer func() {
		if closeError == nil {
			os.Remove(file.Name())
		}
	}()

	if _, err := os.Stat(strings.ToLower(file.Name())); errors.Is(err, os.ErrNotExist) {
		return true, closeError
	}
	return false, closeError
}
