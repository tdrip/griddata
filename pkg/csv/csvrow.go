package csv

import (
	gd "github.com/tdrip/griddata/pkg"
)

//RowData csv implementation of a headed row
type RowData struct {
	gd.RowData
}

//HeadedRowData csv implementation of a headed row
type HeadedRowData struct {
	gd.HeadedRowData
}

//CreateRowData creates a row data from a parsed CSV
func CreateRowData(row int, pass int, datain []string) *RowData {
	rd := gd.CreateRowData(row, pass)
	for d := 0; d < len(datain); d++ {
		cell := gd.CreateStringCell(gd.CreateGDPoint(row, d), datain[d])
		rd.AddCell(cell)
	}

	return &rd
}
