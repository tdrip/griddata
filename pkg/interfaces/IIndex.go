package igrid

const (
	UNKNOWNY = -1
	UNKNOWNX = -1
)

// IIndex This interface has location and relationships
type IIndex interface {
	GetLocation() IPoint
	SetLocation(position IPoint)

	GetRelatedIndexes() []IIndex
	SetRelatedIndexes(relatedindexs []IIndex)
	AddRelatedIndex(relatedi IIndex)
	String() string
}
