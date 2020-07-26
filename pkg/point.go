package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//GDPoint A grid data point
type GDPoint struct {
	igrid.IPoint
	Y int
	X int
}

//CreateGDPoint creates a Grid Data Point
func CreateGDPoint(x int, y int) *GDPoint {
	gdp := GDPoint{Y: y, X: x}
	return &gdp
}

//GetY Gets the Y position
func (point *GDPoint) GetY() int {
	return point.Y
}

//SetY Sets the Y position
func (point *GDPoint) SetY(Y int) {
	point.Y = Y
}

//GetX Gets the X position
func (point *GDPoint) GetX() int {
	return point.X

}

//SetX Sets the X position
func (point *GDPoint) SetX(X int) {
	point.X = X
}

//Match
func (point *GDPoint) Match(position igrid.IPoint) bool {
	return point.Matches(position.GetX(), position.GetY())
}

//Matches This matcches the point based on position
func (point *GDPoint) Matches(X int, Y int) bool {

	// are either negative?
	if X < 0 || Y < 0 {

		if X < 0 && Y < 0 {
			return false
		}

		// we are matching on Y
		if X < 0 {
			return (point.Y == Y)
		}

		// we are matching on X
		if Y < 0 {
			return (point.X == X)
		}
	}

	return (point.X == X) && (point.Y == Y)
}
