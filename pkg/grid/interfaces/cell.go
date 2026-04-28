package igrid

// ICell This interface is data and location of data
type Cell interface {
	GetLocation() Point
	SetLocation(point Point)

	GetData() any
	SetData(data any)
	String() string
}
