package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func GrepSlice(s []string, key string) bool {
	for _, v := range s {
		if strings.Contains(strings.ToLower(v), strings.ToLower(key)) {
			return true
		}
	}

	return false
}

func GrepColor(str string, key string) string {
	index := strings.Index(strings.ToLower(str), strings.ToLower(key))
	colored := str[:index] + color.GreenString(str[index:index+len(key)]) + str[index+len(key):]

	return colored
}

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

func IsDir(path string) bool {
	fi, _ := os.Stat(path)
	return fi.IsDir()
}
