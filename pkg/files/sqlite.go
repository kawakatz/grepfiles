package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/goark/gnkf/guess"
	"github.com/kawakatz/grepfiles/pkg/utils"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/schollz/sqlite3dump"
)

// GrepSqlite3 greps SQLite files.
func GrepSqlite3(path string, keyword string) {
	buf := new(bytes.Buffer)
	_ = sqlite3dump.Dump(path, buf)

	scanner := bufio.NewScanner(strings.NewReader(buf.String()))
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
