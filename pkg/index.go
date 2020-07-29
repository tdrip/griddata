package grid

import (
	"fmt"

	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//Index This struct is the  index
type Index struct {
	igrid.IIndex

	Location igrid.IPoint

	RelatedIndexes []igrid.IIndex
}

//CreateIndex Creates a pointer to the Index struct
func CreateIndex(location igrid.IPoint) *Index {
	gdi := Index{}
	gdi.Location = location
	return &gdi
}

//GetPosition return the position of this Index
func (gdi *Index) GetPosition() igrid.IPoint {
	return gdi.Location
}

//SetPosition set the position of this Index
func (gdi *Index) SetPosition(position igrid.IPoint) {
	gdi.Location = position
}

//GetRelatedIndexes return indexes that are related to this one
func (gdi *Index) GetRelatedIndexes() []igrid.IIndex {
	return gdi.RelatedIndexes
}

//SetRelatedIndexes set indexes that are related to this one
func (gdi *Index) SetRelatedIndexes(relatedi []igrid.IIndex) {
	gdi.RelatedIndexes = relatedi
}

//AddRelatedIndex add an index that is related to this one
func (gdi *Index) AddRelatedIndex(relatedi igrid.IIndex) {
	rindices := gdi.RelatedIndexes
	rindices = append(rindices, relatedi)
	gdi.RelatedIndexes = rindices
}

//String Print Index
func (gdi *Index) String() string {
	return fmt.Sprintf("Index at [%s]", gdi.GetPosition())
}
