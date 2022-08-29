package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/goark/gnkf/guess"
	"github.com/kawakatz/grepfiles/pkg/utils"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

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

		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
