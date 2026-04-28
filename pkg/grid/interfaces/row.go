package igrid

// IRow This interface is data for the row
type Row interface {
	GetIndex() Point
	SetIndex(index Point)

	Matches(index Point) bool

	GetCells() []Cell
	SetCells(cells []Cell)
	AddCell(cell Cell)
}
