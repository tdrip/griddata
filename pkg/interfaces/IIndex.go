package igrid

// IIndex This interface has location
type IIndex interface {
	GetLocation() IPoint
	SetLocation(position IPoint)

	Matches(pos IPoint) bool

	String() string
}
