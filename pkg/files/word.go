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
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	defer f.Close()

	content, _, err := docconv.ConvertDocx(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		ss, err := guess.EncodingBytes([]byte(line))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()+": "+path)
			return
		}
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

// GrepWord1997 greps Word1997 files.
func GrepWord1997(path string, keyword string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	defer f.Close()

	content, _, err := docconv.ConvertDoc(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error()+": "+path)
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		ss, err := guess.EncodingBytes([]byte(line))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()+": "+path)
			return
		}
		enc := ss[0]

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
