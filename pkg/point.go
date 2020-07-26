package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//Point A grid data point
type Point struct {
	igrid.IPoint
	Y int
	X int
}

//CreatePoint creates a Grid Data Point
func CreatePoint(x int, y int) *Point {
	gdp := Point{Y: y, X: x}
	return &gdp
}

//GetY Gets the Y position
func (point *Point) GetY() int {
	return point.Y
}

//SetY Sets the Y position
func (point *Point) SetY(Y int) {
	point.Y = Y
}

//GetX Gets the X position
func (point *Point) GetX() int {
	return point.X
}

//SetX Sets the X position
func (point *Point) SetX(X int) {
	point.X = X
}

//Match This matches one point against this one
func (point *Point) Match(position igrid.IPoint) bool {
	return point.Matches(position.GetX(), position.GetY())
}

//Matches This matcches the point based on position
func (point *Point) Matches(X int, Y int) bool {

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

//String Prints the point as X:,Y:
func (point Point) String() string {
	return fmt.Sprintf("X:%d,Y:%d", point.GetX(), point.GetY())
}
