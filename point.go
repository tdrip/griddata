package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

type GDPoint struct {
	igrid.IPoint
	Column int
	Row    int
}

//GetColumn Gets the column position
func (point *GDPoint) GetColumn() int {
	return point.Column
}

//SetColumn Sets the column position
func (point *GDPoint) SetColumn(column int) {
	point.Column = column
}

//GetRow Gets the row position
func (point *GDPoint) GetRow() int {
	return point.Row

}

//SetRow Sets the row position
func (point *GDPoint) SetRow(row int) {
	point.Row = row
}

//Match
func (point *GDPoint) Match(position igrid.IPoint) bool {
	return point.Matches(position.GetRow(), position.GetColumn())
}

//Matches This matcches the point based on position
func (point *GDPoint) Matches(row int, column int) bool {

	// are either negative?
	if row < 0 || column < 0 {

		if column < 0 && row < 0 {
			return false
		}

		// we are matching on row
		if column < 0 {
			return (point.Row == row)
		}

		// we are matching on column
		if row < 0 {
			return (point.Column == column)
		}
	}

	return (point.Column == column) && (point.Row == row)
}
