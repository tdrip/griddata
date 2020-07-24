package igrid

//IColumn This interface is data for the column
type IColumn interface {
	GetIndex() IIndex
	SetIndex(index IIndex)

	Matches(index IIndex) bool

	GetCells() []ICell
	SetCells(cells []ICell)
}

//IHeadedColumn This column has a headers
type IHeadedColumn interface {
	IColumn
	GetHeader() IHeader
	SetHeader(header IHeader)
}
