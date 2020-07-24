package igrid

//ICell This interface is data and location of data
type ICell interface {
	GetLocation() IPoint
	SetLocation(point IPoint)

	GetData() interface{}
	SetData(data interface{})
}
