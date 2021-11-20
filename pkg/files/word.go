package files

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"code.sajari.com/docconv"

	"github.com/kawakatz/grepfiles/pkg/utils"
)

func GrepWord2007(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertDocx(f)
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			pwd, _ := os.Getwd()
			rel, _ := filepath.Rel(pwd, path)
			fmt.Print(rel, ":")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}

func GrepWord1997(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertDoc(f)
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			pwd, _ := os.Getwd()
			rel, _ := filepath.Rel(pwd, path)
			fmt.Print(rel, ":")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
