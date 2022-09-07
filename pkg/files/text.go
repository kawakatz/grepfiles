package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

	"github.com/goark/gnkf/guess"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// GrepText greps text files.
func GrepText(path string, keyword string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	defer f.Close()

	s, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	ss, err := guess.EncodingBytes(s)
	var enc string
	if err == nil {
		enc = ss[0]
	} else {
		enc = ""
	}

	if enc == "UTF-16LE" {
		reader := bufio.NewReader(transform.NewReader(f, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
		contents, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()+": "+path)
			return
		}
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
				u8, err := ioutil.ReadAll(transform.NewReader(reader, japanese.ISO2022JP.NewDecoder()))
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error()+": "+path)
					return
				}
				line = string(u8)
			case "EUCJP":
				reader := strings.NewReader(line)
				u8, err := ioutil.ReadAll(transform.NewReader(reader, japanese.EUCJP.NewDecoder()))
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error()+": "+path)
					return
				}
				line = string(u8)
			case "Shift_JIS", "windows-1252":
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
}
