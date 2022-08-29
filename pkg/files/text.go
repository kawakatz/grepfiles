package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/goark/gnkf/guess"
	"github.com/kawakatz/grepfiles/pkg/utils"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// GrepText greps text files.
func GrepText(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	s, _ := ioutil.ReadFile(path)
	ss, _ := guess.EncodingBytes(s)
	enc := ss[0]

	if enc == "UTF-16LE" {
		reader := bufio.NewReader(transform.NewReader(f, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
		contents, _ := ioutil.ReadAll(reader)
		enter := "\r\n|\n"
		array := regexp.MustCompile(enter).Split(string(contents), -1)
		for _, line := range array {
			if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
				fmt.Print(path, ": ")
				fmt.Println(utils.GrepColor(line, keyword))
			}
		}
	} else {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
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
}
