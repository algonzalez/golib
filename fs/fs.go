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

// GetFileInfo returns an fs.FileInfo representing the file at path.
// If the file was not found, it will return a custom error.
// If there is another error, it will be of type *fs.PathError
func GetFileInfo(path string) (fs.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, errors.New("requested path is a directory not file")
	}
	return info, err
}

// GetDirInfo returns an fs.FileInfo representing the directory at path.
// If the file was not found, it will return a custom error.
// If there is another error, it will be of type *fs.PathError
func GetDirInfo(path string) (fs.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("requested path is not a directory")
	}
	return info, err
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
