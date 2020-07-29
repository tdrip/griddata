package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//RowData This represents a row of data
type RowData struct {
	igrid.IRow

	// Index of the row
	Index igrid.IIndex

	// Number of passes over the row
	Pass int

	// Raw data
	RawData []interface{}

	// Parsed Cell Data
	Cells []igrid.ICell
}

//HeadedRowData a row with a header
type HeadedRowData struct {
	igrid.IHeadedRow

	Header igrid.IHeader
}

//CreateRowData Creates a default tow data struct
func CreateRowData(row int, pass int) *RowData {
	rd := RowData{Pass: pass}

	// x,y point
	rowp := CreatePoint(row, -1)

	// set the index
	rd.SetIndex(CreateIndex(rowp))

	return &rd
}

//GetIndex Gets the index for the row
func (rd *RowData) GetIndex() igrid.IIndex {
	return rd.Index
}

//SetIndex Sets the index for the row
func (rd *RowData) SetIndex(index igrid.IIndex) {
	rd.Index = index
}

//Matches Matches the index passed in against the index for the row
func (rd *RowData) Matches(index igrid.IIndex) bool {
	return rd.GetIndex().GetLocation().Match(index.GetLocation())
}

//GetCells Gets the cells for the row
func (rd *RowData) GetCells() []igrid.ICell {
	return rd.Cells
}

//SetCells Sets the cells for the row
func (rd *RowData) SetCells(cells []igrid.ICell) {
	rd.Cells = cells
}

//AddCell Add a cells to the row
func (rd *RowData) AddCell(cell igrid.ICell) {
	cells := rd.Cells
	cells = append(cells, cell)
	rd.Cells = cells
}

//GetHeader Returns the header for the row
func (rd *HeadedRowData) GetHeader() igrid.IHeader {
	return rd.Header
}

//SetHeader Sets the header for the row
func (rd *HeadedRowData) SetHeader(header igrid.IHeader) {
	rd.Header = header
}
