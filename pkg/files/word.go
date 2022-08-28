package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"code.sajari.com/docconv"

	"github.com/kawakatz/grepfiles/pkg/utils"
)

// GrepWord2007 greps Word2007 files.
func GrepWord2007(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertDocx(f)
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}

// GrepWord1997 greps Word1997 files.
func GrepWord1997(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertDoc(f)
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
