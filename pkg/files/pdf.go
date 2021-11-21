package files

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

	"code.sajari.com/docconv"
)

// GrepPDF greps PDF files.
func GrepPDF(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertPDF(f)
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
