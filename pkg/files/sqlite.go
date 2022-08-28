package files

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

	"github.com/schollz/sqlite3dump"
)

// GrepSqlite3 greps SQLite files.
func GrepSqlite3(path string, keyword string) {
	buf := new(bytes.Buffer)
	_ = sqlite3dump.Dump(path, buf)

	scanner := bufio.NewScanner(strings.NewReader(buf.String()))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			fmt.Print(path, ": ")
			fmt.Println(utils.GrepColor(line, keyword))
		}
	}
}
