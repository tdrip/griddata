package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//Index This struct is the  index
type Index struct {
	igrid.IIndex

	Location igrid.IPoint

	RelatedIndexes []igrid.IIndex
}

func CreateIndex(location igrid.IPoint) *Index {
	gdi := Index{}
	gdi.Location = location
	return &gdi
}

func (gdi *Index) GetPosition() igrid.IPoint {
	return gdi.Location
}

func (gdi *Index) SetPosition(position igrid.IPoint) {
	gdi.Location = position
}

func (gdi *Index) GetRelatedIndexes() []igrid.IIndex {
	return gdi.RelatedIndexes
}

func (gdi *Index) SetRelatedIndexes(relatedi []igrid.IIndex) {
	gdi.RelatedIndexes = relatedi
}

func (gdi *Index) AddRelatedIndex(relatedi igrid.IIndex) {
	rindices := gdi.RelatedIndexes
	rindices = append(rindices, relatedi)
	gdi.RelatedIndexes = rindices
}
