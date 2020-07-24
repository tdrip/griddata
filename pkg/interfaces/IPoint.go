package igrid

//IPoint This represents where the Cell is
type IPoint interface {
	GetX() int
	SetX(X int)

	GetY() int
	SetY(Y int)

	Matches(Y int, X int) bool
	Match(point IPoint) bool
}
