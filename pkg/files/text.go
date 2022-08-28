package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"
)

// GrepText greps text files.
func GrepText(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
