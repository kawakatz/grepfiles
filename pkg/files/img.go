package files

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

	gosseract "github.com/otiai10/gosseract/v2"
)

// GrepImg greps image files.
func GrepImg(path string, keyword string) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("eng", "jpn")
	client.SetImage(path)

	text, _ := client.Text()
	scanner := bufio.NewScanner(strings.NewReader(text))
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
