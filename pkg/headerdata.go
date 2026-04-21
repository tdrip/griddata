package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// RowData This represents a row of data
type HeaderRowData struct {
	igrid.IRow

	// Index of the row
	Index igrid.IIndex

	// Number of passes over the row
	Pass int

	// Raw data
	RawData []any

	// Parsed Cell Data
	Cells []igrid.ICell

	IndexedRawData map[int]any
}

// CreateRowData Creates a default row data struct
func CreateHeaderRowData(row int, pass int) *HeaderRowData {
	rd := HeaderRowData{Pass: pass}

	// x,y point
	rowp := CreatePoint(row, -1)

	// set the index
	rd.SetIndex(CreateIndex(rowp))

	return &rd
}

// GetIndex Gets the index for the row
func (rd *HeaderRowData) GetIndex() igrid.IIndex {
	return rd.Index
}

// SetIndex Sets the index for the row
func (rd *HeaderRowData) SetIndex(index igrid.IIndex) {
	rd.Index = index
}

// Matches Matches the index passed in against the index for the row
func (rd *HeaderRowData) Matches(index igrid.IIndex) bool {
	return rd.GetIndex().GetLocation().Match(index.GetLocation())
}

// GetCells Gets the cells for the row
func (rd *HeaderRowData) GetCells() []igrid.ICell {
	return rd.Cells
}

// SetCells Sets the cells for the row
func (rd *HeaderRowData) SetCells(cells []igrid.ICell) {
	rd.Cells = cells
}

// AddCell Add a cells to the row
func (rd *HeaderRowData) AddCell(cell igrid.ICell) {
	cells := rd.Cells
	cells = append(cells, cell)
	rd.Cells = cells
}

// FillHeaderRowStringData creates a row data from a string data array
func FillHeaderRowStringData(rowindex int, pass int, columndata []string) *HeaderRowData {

	// number of passes and the row index
	rd := CreateHeaderRowData(rowindex, pass)
	indexeddata := make(map[int]any, len(columndata))
	for columnindex := 0; columnindex < len(columndata); columnindex++ {

		pnt := CreatePoint(rowindex, columnindex)
		// csv is always srting so we parse the cells as such
		cell := CreateStringCell(pnt, columndata[columnindex])

		indexeddata[columnindex] = columndata[columnindex]

		// add the cell
		rd.AddCell(cell)
	}

	rd.IndexedRawData = indexeddata

	return rd
}
