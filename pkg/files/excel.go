package files

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/goark/gnkf/guess"
	"github.com/kawakatz/grepfiles/pkg/utils"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/extrame/xls"
	excelize "github.com/xuri/excelize/v2"
)

// GrepExcel2007 greps Excel2007 files.
func GrepExcel2007(path string, keyword string) {
	f, _ := excelize.OpenFile(path)
	sheetNames := f.GetSheetList()

	for _, sheetName := range sheetNames {
		rows, _ := f.GetRows(sheetName)
		for _, row := range rows {
			if utils.GrepSlice(row, keyword) {
				fmt.Print(path, ": ")
				outStr := strings.Join(row, ",")

				ss, _ := guess.EncodingBytes([]byte(outStr))
				enc := ss[0]

				switch enc {
				case "ISO2022JP":
					reader := strings.NewReader(outStr)
					u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.ISO2022JP.NewDecoder()))
					outStr = string(u8)
				case "EUCJP":
					reader := strings.NewReader(outStr)
					u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.EUCJP.NewDecoder()))
					outStr = string(u8)
				case "Shift_JIS":
					reader := strings.NewReader(outStr)
					u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.ShiftJIS.NewDecoder()))
					outStr = string(u8)
				default:
					break
				}

				fmt.Println(utils.GrepColor(outStr, keyword))
			}
		}
	}
}

// GrepExcel1997 greps Excel1997 files.
func GrepExcel1997(path string, keyword string) {
	f, _ := xls.Open(path, "utf-8")

	sheetNum := f.NumSheets()
	for i := 0; i < sheetNum; i++ {
		sheet := f.GetSheet(i)
		if sheet != nil {
			for row := 0; row < int(sheet.MaxRow); row++ {
				r := sheet.Row(row)
				rowSlice := []string{}
				for col := r.FirstCol(); col < r.LastCol(); col++ {
					value := r.Col(col)
					rowSlice = append(rowSlice, value)
				}
				if utils.GrepSlice(rowSlice, keyword) {
					fmt.Print(path, ": ")
					outStr := strings.Join(rowSlice, ",")

					ss, _ := guess.EncodingBytes([]byte(outStr))
					enc := ss[0]

					switch enc {
					case "ISO2022JP":
						reader := strings.NewReader(outStr)
						u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.ISO2022JP.NewDecoder()))
						outStr = string(u8)
					case "EUCJP":
						reader := strings.NewReader(outStr)
						u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.EUCJP.NewDecoder()))
						outStr = string(u8)
					case "Shift_JIS":
						reader := strings.NewReader(outStr)
						u8, _ := ioutil.ReadAll(transform.NewReader(reader, japanese.ShiftJIS.NewDecoder()))
						outStr = string(u8)
					default:
						break
					}

					fmt.Println(utils.GrepColor(outStr, keyword))
				}
			}
		}
	}
}
