package igrid

//IPoint This represents where the Cell is
type IPoint interface {
	GetColumn() int
	SetColumn(column int)

	GetRow() int
	SetRow(row int)

	Matches(row int, column int) bool
	Match(point IPoint) bool
}
