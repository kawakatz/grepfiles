package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kawakatz/grepfiles/pkg/utils"

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
				pwd, _ := os.Getwd()
				rel, _ := filepath.Rel(pwd, path)
				fmt.Print(rel, ":")
				outStr := strings.Join(row, ",")
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
					pwd, _ := os.Getwd()
					rel, _ := filepath.Rel(pwd, path)
					fmt.Print(rel, ":")
					outStr := strings.Join(rowSlice, ",")
					fmt.Println(utils.GrepColor(outStr, keyword))
				}
			}
		}
	}
}
