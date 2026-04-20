package igrid

// IRow This interface is data for the row
type IRow interface {
	GetIndex() IIndex
	SetIndex(index IIndex)

	Matches(index IIndex) bool

	GetCells() []ICell
	SetCells(cells []ICell)
	AddCell(cell ICell)
}
