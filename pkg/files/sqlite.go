package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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
		ss, err := guess.EncodingBytes([]byte(line))
		var enc string
		if err == nil {
			enc = ss[0]
		} else {
			enc = ""
		}

		switch {
		case enc == "ISO2022JP":
			reader := strings.NewReader(line)
			u8, err := ioutil.ReadAll(transform.NewReader(reader, japanese.ISO2022JP.NewDecoder()))
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error()+": "+path)
				return
			}
			line = string(u8)
		case enc == "EUCJP":
			reader := strings.NewReader(line)
			u8, err := ioutil.ReadAll(transform.NewReader(reader, japanese.EUCJP.NewDecoder()))
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error()+": "+path)
				return
			}
			line = string(u8)
		case enc == "Shift_JIS", strings.HasPrefix(enc, "windows-"):
			reader := strings.NewReader(line)
			u8, err := ioutil.ReadAll(transform.NewReader(reader, japanese.ShiftJIS.NewDecoder()))
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error()+": "+path)
				return
			}
			line = string(u8)
		}

		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
