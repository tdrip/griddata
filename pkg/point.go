package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//Point A grid data point
type Point struct {
	igrid.IPoint
	y int
	x int
}

//CreatePoint creates a Grid Data Point
func CreatePoint(x int, y int) *Point {
	gdp := Point{y: y, x: x}
	return &gdp
}

//GetY Gets the Y position
func (point *Point) GetY() int {
	return point.y
}

//SetY Sets the Y position
func (point *Point) SetY(Y int) {
	point.y = Y
}

//GetX Gets the X position
func (point *Point) GetX() int {
	return point.x
}

//SetX Sets the X position
func (point *Point) SetX(X int) {
	point.x = X
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
			return (point.y == Y)
		}

		// we are matching on X
		if Y < 0 {
			return (point.x == X)
		}
	}

	return (point.x == X) && (point.y == Y)
}

//String Prints the point as X:,Y:
func (point Point) String() string {
	return fmt.Sprintf("X:%d,Y:%d", point.GetX(), point.GetY())
}
