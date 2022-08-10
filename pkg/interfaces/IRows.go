package igrid

//IRow This interface is data for the row
type IRow interface {
	GetIndex() IIndex
	SetIndex(index IIndex)

	Matches(index IIndex) bool

	GetCells() []ICell
	SetCells(cells []ICell)
	AddCell(cell ICell)
}

//IHeadedRow This row has a headers
// <SOME HEADER> and all cells after this
type IHeadedRow interface {
	IRow
	GetHeader() IHeader
	SetHeader(header IHeader)
}

/*
type IRowAction interface {
	IDataAction
	PeformAction(IRow)
}
*/
