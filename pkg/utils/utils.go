package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// Usage prints usage.
func Usage() {
	fmt.Println("usage: grepfiles <path> <keyword>")
}

// GrepSlice checks if the keyword is in the slice.
func GrepSlice(s []string, key string) bool {
	for _, v := range s {
		if strings.Contains(strings.ToLower(v), strings.ToLower(key)) {
			return true
		}
	}

	return false
}

// GrepColor colors the keyword section.
func GrepColor(str string, key string) string {
	index := strings.Index(strings.ToLower(str), strings.ToLower(key))
	colored := str[:index] + color.GreenString(str[index:index+len(key)]) + str[index+len(key):]

	return colored
}

// LsR recursively lists the directory.
func LsR(dir string) []string {
	files, _ := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, LsR(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

// IsDir checks if the path is a directory.
func IsDir(path string) bool {
	fi, _ := os.Stat(path)
	return fi.IsDir()
}

// IsExist checks if the path exists.
func IsExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
