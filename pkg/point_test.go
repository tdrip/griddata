package grid

import (
	"testing"
)

func TestPointMatches(t *testing.T) {

	firstpoint := CreatePoint(1, 1)

	if firstpoint.GetX() != 1 || firstpoint.GetY() != 1 {
		t.Errorf("X was %d not 1 - Y was %d not 1", firstpoint.GetX(), firstpoint.GetY())
	}

	if !firstpoint.Matches(1, 1) {
		t.Errorf("%s did not match 1,1 ", firstpoint)
	}
	//t.Errorf("%s", firstpoint)
}

func TestPointMatch(t *testing.T) {

	firstpoint := CreatePoint(1, 1)

	if firstpoint.GetX() != 1 || firstpoint.GetY() != 1 {
		t.Errorf("X was %d not 1 - Y was %d not 1", firstpoint.GetX(), firstpoint.GetY())
	}

	secondpoint := CreatePoint(1, 1)

	if !firstpoint.Match(secondpoint) {
		t.Errorf("%s did not match %s ", firstpoint, secondpoint)
	}

}

func TestPointSet(t *testing.T) {

	firstpoint := CreatePoint(1, 1)

	if firstpoint.GetX() != 1 || firstpoint.GetY() != 1 {
		t.Errorf("X was %d not 1 - Y was %d not 1", firstpoint.GetX(), firstpoint.GetY())
	}

	firstpoint.SetX(2)

	if firstpoint.GetX() != 2 {
		t.Errorf("X was %d not 2 ", firstpoint.GetX())
	}

	firstpoint.SetY(2)

	if firstpoint.GetY() != 2 {
		t.Errorf("Y was %d not 2 ", firstpoint.GetY())
	}
}
