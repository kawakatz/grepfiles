package grep

import (
	"io/ioutil"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/files"

	"github.com/h2non/filetype"
)

// GrepFile greps depends on the file formats.
func GrepFile(path string, keyword string) {
	buf, _ := ioutil.ReadFile(path)
	kind, _ := filetype.Match(buf)
	mime := kind.MIME.Value

	if strings.Contains(mime, "image") {
		files.GrepImg(path, keyword)
	} else if mime == "application/pdf" {
		files.GrepPDF(path, keyword)
	} else if mime == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		files.GrepExcel2007(path, keyword)
	} else if mime == "application/vnd.ms-excel" {
		//files.GrepExcel1997(path, keyword)
	} else if mime == "application/vnd.openxmlformats-officedocument.presentationml.presentation" {
		files.GrepPowerPoint2007(path, keyword)
	} else if mime == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
		files.GrepWord2007(path, keyword)
	} else if mime == "application/msword" {
		files.GrepWord1997(path, keyword)
	} else if mime == "application/vnd.sqlite3" {
		files.GrepSqlite3(path, keyword)
	} else {
		files.GrepText(path, keyword)
	}
}
