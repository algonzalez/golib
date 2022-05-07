package fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/algonzalez/golib/path"
	"github.com/algonzalez/golib/slice"
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

// GetSearchPaths returns a distinct list of paths
// defined in the PATH environment value.
// Also includes the current working directory in the list.
func GetSearchPaths() []string {
	if pathString, ok := os.LookupEnv("PATH"); ok {
		paths := filepath.SplitList(pathString)

		// include current directory
		cwd, _ := os.Getwd()
		paths = slice.Prepend(paths, cwd)

		return slice.ToDistinct(paths, path.NormalizePath)
	}
	return []string{}
}
