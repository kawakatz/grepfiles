package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"code.sajari.com/docconv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/goark/gnkf/guess"
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

// GrepWord1997 greps Word1997 files.
func GrepWord1997(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertDoc(f)
	scanner := bufio.NewScanner(strings.NewReader(content))
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
