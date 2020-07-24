package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//GDRowData
type GDRowData struct {
	igrid.IRow

	Index igrid.IIndex

	// raw data
	RawData []interface{}

	// mapped data
	Pass int

	// Parsed Cell Data
	Cells []igrid.ICell
}

type GDHeadedRowData struct {
	igrid.IHeadedRow

	Header igrid.IHeader
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

func (rd *GDHeadedRowData) GetHeader() igrid.IHeader {
	return rd.Header
}

func (rd *GDHeadedRowData) SetHeader(header igrid.IHeader) {
	rd.Header = header
}
