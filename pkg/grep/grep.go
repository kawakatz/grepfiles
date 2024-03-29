package grep

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/files"

	"github.com/h2non/filetype"
)

// GrepFile greps depends on the file formats.
func GrepFile(path string, keyword string) {
	buf, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	kind, err := filetype.Match(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	mime := kind.MIME.Value
	//fmt.Println(path + ": " + mime)

	switch mime {
	case "application/zip":
		return
	case "application/x-tar":
		return
	case "application/x-7z-compressed":
		return
	case "application/vnd.rar":
		return
	case "application/gzip":
		return
	case "application/x-bzip2":
		return
	case "application/x-lzip":
		return
	case "application/x-xz":
		return
	}

	if strings.Contains(mime, "image") {
		files.GrepImg(path, keyword)
	} else if mime == "application/pdf" || filepath.Ext(path) == ".pdf" {
		files.GrepPDF(path, keyword)
	} else if mime == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" || filepath.Ext(path) == ".xlsx" || filepath.Ext(path) == ".xlsm" {
		files.GrepExcel2007(path, keyword)
	} else if mime == "application/vnd.ms-excel" || filepath.Ext(path) == ".xls" {
		files.GrepExcel1997(path, keyword)
	} else if mime == "application/vnd.openxmlformats-officedocument.presentationml.presentation" || filepath.Ext(path) == ".pptx" {
		files.GrepPowerPoint2007(path, keyword)
	} else if mime == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" || filepath.Ext(path) == ".docx" || filepath.Ext(path) == ".docm" {
		files.GrepWord2007(path, keyword)
	} else if mime == "application/msword" || filepath.Ext(path) == ".doc" {
		files.GrepWord1997(path, keyword)
	} else if mime == "application/vnd.sqlite3" {
		files.GrepSqlite3(path, keyword)
	} else {
		files.GrepText(path, keyword)
	}
}
