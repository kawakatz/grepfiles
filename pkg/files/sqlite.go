package files

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

	"github.com/schollz/sqlite3dump"
)

func GrepSqlite3(path string, keyword string) {
	buf := new(bytes.Buffer)
	_ = sqlite3dump.Dump(path, buf)

	scanner := bufio.NewScanner(strings.NewReader(buf.String()))
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
