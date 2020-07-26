package grid

import (
	igrid "github.com/tdrip/griddata/pkg/interfaces"
)

//GDIndex This struct is the  index
type GDIndex struct {
	igrid.IIndex

	Location igrid.IPoint

	RelatedIndexes []igrid.IIndex
}

func CreateGDIndex(location igrid.IPoint) *GDIndex {
	gdi := GDIndex{}
	gdi.Location = location
	return &gdi
}

func (gdi *GDIndex) GetPosition() igrid.IPoint {
	return gdi.Location
}

func (gdi *GDIndex) SetPosition(position igrid.IPoint) {
	gdi.Location = position
}

func (gdi *GDIndex) GetRelatedIndexes() []igrid.IIndex {
	return gdi.RelatedIndexes
}

func (gdi *GDIndex) SetRelatedIndexes(relatedi []igrid.IIndex) {
	gdi.RelatedIndexes = relatedi
}

func (gdi *GDIndex) AddRelatedIndex(relatedi igrid.IIndex) {
	rindices := gdi.RelatedIndexes
	rindices = append(rindices, relatedi)
	gdi.RelatedIndexes = rindices
}
