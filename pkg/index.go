package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

// Index This struct is the  index
type Index struct {
	igrid.IIndex

	Location igrid.IPoint
}

// JustXIndex ceates a pointer to the Index struct with just x set
func JustXIndex(x int) *Index {
	gdi := Index{}
	gdi.Location = NewPoint(x, igrid.UNKNOWNY)
	return &gdi
}

// JustYIndex ceates a pointer to the Index struct with just x set
func JustYIndex(y int) *Index {
	gdi := Index{}
	gdi.Location = NewPoint(igrid.UNKNOWNX, y)
	return &gdi
}

// NewIndex ceates a pointer to the Index struct
func NewIndex(location igrid.IPoint) *Index {
	gdi := Index{}
	gdi.Location = location
	return &gdi
}

// GetPosition return the position of this Index
func (gdi *Index) GetLocation() igrid.IPoint {
	return gdi.Location
}

// SetPosition set the position of this Index
func (gdi *Index) SetLocation(position igrid.IPoint) {
	gdi.Location = position
}

// String Print Index
func (gdi *Index) String() string {
	return fmt.Sprintf("Index at [%s]", gdi.GetLocation())
}

// Does it match a point?
func (gdi *Index) Matches(pos igrid.IPoint) bool {
	return gdi.GetLocation().Match(pos)
}
