package grid

import (
	igrid "github.com/tdrip/griddata/pkg/grid/interfaces"
)

// Row This represents a row of data
type Row struct {
	igrid.IRow

	// Index of the row
	Index igrid.IPoint

	// Number of passes over the row
	Pass int

	// Parsed Cell Data
	// index these!
	Cells []igrid.ICell
}

// NewRow creates a default row data struct
func NewRow(rowindex int, pass int) *Row {
	rd := Row{Pass: pass}

	// x,y point doesn't matter
	// just need X as this is a row
	// set the index
	rd.SetIndex(JustXPoint(rowindex))

	return &rd
}

// GetIndex Gets the index for the row
func (rd *Row) GetIndex() igrid.IPoint {
	return rd.Index
}

// SetIndex Sets the index for the row
func (rd *Row) SetIndex(index igrid.IPoint) {
	rd.Index = index
}

// Matches Matches the index passed in against the index for the row
func (rd *Row) Matches(index igrid.IPoint) bool {
	return rd.GetIndex().Match(index)
}

// GetCells Gets the cells for the row
func (rd *Row) GetCells() []igrid.ICell {
	return rd.Cells
}

// SetCells Sets the cells for the row
func (rd *Row) SetCells(cells []igrid.ICell) {
	rd.Cells = cells
}

// AddCell Add a cells to the row
func (rd *Row) AddCell(cell igrid.ICell) {
	cells := rd.Cells
	cells = append(cells, cell)
	rd.Cells = cells
}
