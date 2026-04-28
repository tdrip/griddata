package igrid

// IRow This interface is data for the row
type IRow interface {
	GetIndex() IPoint
	SetIndex(index IPoint)

	Matches(index IPoint) bool

	GetCells() []ICell
	SetCells(cells []ICell)
	AddCell(cell ICell)
}
