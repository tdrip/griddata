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
func CreateRowData(rowindex int, pass int, columndata []string) *gd.RowData {

	// number of passes and the row index
	rd := gd.CreateRowData(rowindex, pass)

	for columnindex := 0; columnindex < len(columndata); columnindex++ {

		pnt := gd.CreatePoint(rowindex, columnindex)
		// csv is always srting so we parse the cells as such
		cell := gd.CreateStringCell(pnt, columndata[columnindex])

		// add the cell
		rd.AddCell(cell)
	}

	return rd
}
