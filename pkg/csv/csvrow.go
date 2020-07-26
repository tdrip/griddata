package csv

import (
	gd "github.com/tdrip/griddata/pkg"
)

//RowData csv implementation of a headed row
type RowData struct {
	gd.GDRowData
}

//HeadedRowData csv implementation of a headed row
type HeadedRowData struct {
	gd.GDHeadedRowData
}

//CreateRowData
func CreateRowData(row int, datain []string) *RowData {
	rd := RowData{}

	cells := []*gd.GDCell{}
	//rd.SetIndex
	for d := 0; d < len(datain); d++ {
		cell := gd.CreateStringCell(gd.CreateGDPoint(row, d), datain[d])
		cells := append(cells)
	}

	return &rd
}
