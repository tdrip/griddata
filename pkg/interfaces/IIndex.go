package igrid

//IIndex This interface has location and relationships
type IIndex interface {
	GetLocation() IPoint
	SetLocation(position IPoint)

	GetRelatedIndexes() []IIndex
	SetRelatedIndexes(columns []IIndex)
}

//IHeader This interface can be a rowheader, column header etc
type IHeader interface {
	IIndex

	GetDisplayName() string
	SetDisplayName(displayname string)

	Matches() bool
}
