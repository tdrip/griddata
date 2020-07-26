package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//GDRowData
type GDRowData struct {
	igrid.IRow

	// Index of the row
	Index igrid.IIndex

	// mapped data
	Pass int

	// raw data
	RawData []interface{}

	// Parsed Cell Data
	Cells []igrid.ICell
}

type GDHeadedRowData struct {
	igrid.IHeadedRow

	Header igrid.IHeader
}

//CreateRowData
func CreateGDRowData(row int, pass int) *GDRowData {
	rd := GDRowData{}
	rowp := CreateGDPoint(row, -1)
	rd.SetIndex(CreateGDIndex(rowp))
	return &rd
}

func (rd *GDRowData) GetIndex() igrid.IIndex {
	return rd.Index
}

func (rd *GDRowData) SetIndex(index igrid.IIndex) {
	rd.Index = index
}

func (rd *GDRowData) Matches(index igrid.IIndex) bool {
	return rd.GetIndex().GetLocation().Match(index.GetLocation())
}

func (rd *GDRowData) GetCells() []igrid.ICell {
	return rd.Cells
}

func (rd *GDRowData) SetCells(cells []igrid.ICell) {
	rd.Cells = cells
}

//AddCell Add a cells to the row
func (rd *GDRowData) AddCell(cell igrid.ICell) {
	cells := rd.Cells
	cells = append(cells, cell)
	rd.Cells = cells
}

func (rd *GDHeadedRowData) GetHeader() igrid.IHeader {
	return rd.Header
}

func (rd *GDHeadedRowData) SetHeader(header igrid.IHeader) {
	rd.Header = header
}
