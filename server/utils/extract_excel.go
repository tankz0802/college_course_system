package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func extractExcel(path string) {
	xls, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xls.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xls.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

