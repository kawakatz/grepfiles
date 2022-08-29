package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/goark/gnkf/guess"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Usage prints usage.
func Usage() {
	fmt.Println("usage: grepfiles <path> <keyword>")
}

// GrepSlice checks if the keyword is in the slice.
func GrepSlice(s []string, key string) bool {
	for _, line := range s {
		ss, _ := guess.EncodingBytes([]byte(line))
		enc := ss[0]

		switch enc {
		case "ISO2022JP":
			reader := strings.NewReader(line)
			u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.ISO2022JP.NewDecoder()))
			line = string(u8)
		case "EUCJP":
			reader := strings.NewReader(line)
			u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.EUCJP.NewDecoder()))
			line = string(u8)
		case "Shift_JIS":
			reader := strings.NewReader(line)
			u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.ShiftJIS.NewDecoder()))
			line = string(u8)
		default:
			break
		}

		if strings.Contains(strings.ToLower(line), strings.ToLower(key)) {
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
