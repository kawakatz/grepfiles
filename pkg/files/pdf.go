package files

import (
	"bufio"
	"fmt"
	"os"
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
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
