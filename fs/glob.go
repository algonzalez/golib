package fs

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Glob returns the names of all files and directories matching pattern
// or nil if there are no matches.
//
// It uses os.ReadDir and thus may return any error returned from that call.
func Glob(path string, pattern string, ignoreCase bool) (matches []string, err error) {
	return glob(path, pattern, ignoreCase, func(entry fs.DirEntry) bool { return true })
}

// GlobDirs returns the names of all directories matching pattern
// or nil if there are no matches.
//
// It uses os.ReadDir and thus may return any error returned from that call.
func GlobDirs(path string, pattern string, ignoreCase bool) (matches []string, err error) {
	return glob(path, pattern, ignoreCase, func(entry fs.DirEntry) bool { return entry.IsDir() })
}

// GlobFiles returns the names of all files and directories matching pattern
// or nil if there are no matches.
//
// It uses os.ReadDir and thus may return any error returned from that call.
func GlobFiles(path string, pattern string, ignoreCase bool) (matches []string, err error) {
	return glob(path, pattern, ignoreCase, func(entry fs.DirEntry) bool { return !entry.IsDir() })
}

func glob(path string, pattern string, ignoreCase bool, filter func(entry fs.DirEntry) bool) (matches []string, err error) {
	pattern = globToRegex(pattern, ignoreCase)
	var dirEntries []fs.DirEntry
	dirEntries, err = os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range dirEntries {
		if filter(entry) {
			matched, _ := regexp.MatchString(pattern, entry.Name())
			if matched {
				matches = append(matches, filepath.Join(path, entry.Name()))
			}
		}
	}
	return matches, err
}

func globToRegex(pattern string, ignoreCase bool) string {
	rePattern := strings.ReplaceAll(pattern, `\`, `\\`)  // escape
	rePattern = strings.ReplaceAll(rePattern, `.`, `\.`) // escape periods
	rePattern = strings.ReplaceAll(rePattern, `*`, `.*`) // 0 or many of any character
	rePattern = strings.ReplaceAll(rePattern, `?`, `.?`) // 0 or 1 of any character
	if ignoreCase {
		rePattern = `(?i)` + rePattern
	}
	return `^` + rePattern + `$` // full string match
}
