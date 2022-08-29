package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/goark/gnkf/guess"
	"github.com/kawakatz/pkg/utils"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"code.sajari.com/docconv"
)

// GrepPowerPoint2007 greps PowerPoint2007 files.
func GrepPowerPoint2007(path string, keyword string) {
	f, _ := os.Open(path)
	defer f.Close()

	content, _, _ := docconv.ConvertPptx(f)
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
