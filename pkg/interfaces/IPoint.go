package igrid

//IPoint This represents where the Cell is
type IPoint interface {
	// Gets the X Co-Ordinate where it is in the GRID
	GetX() int
	SetX(X int)

	// Gets the Y Co-Ordinate where it is in the GRID
	GetY() int
	SetY(Y int)

	// Do matches
	Matches(Y int, X int) bool
	Match(point IPoint) bool
}
