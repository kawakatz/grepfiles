package files

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

	"github.com/otiai10/gosseract/v2"
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
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
