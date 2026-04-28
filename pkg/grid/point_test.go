package grid

import (
	"testing"
)

func TestPointMatches(t *testing.T) {

	firstpoint := NewPoint(1, 1)

	if firstpoint.GetX() != 1 || firstpoint.GetY() != 1 {
		t.Errorf("X was %d not 1 - Y was %d not 1", firstpoint.GetX(), firstpoint.GetY())
	}

	t.Logf("X was %d - Y was %d", firstpoint.GetX(), firstpoint.GetY())

	if !firstpoint.Matches(1, 1) {
		t.Errorf("%s did not match 1,1 ", firstpoint)
	}
	//t.Errorf("%s", firstpoint)
}

func TestPointMatch(t *testing.T) {

	firstpoint := NewPoint(1, 1)

	if firstpoint.GetX() != 1 || firstpoint.GetY() != 1 {
		t.Errorf("X was %d not 1 - Y was %d not 1", firstpoint.GetX(), firstpoint.GetY())
	}

	secondpoint := NewPoint(1, 1)

	if !firstpoint.Match(secondpoint) {
		t.Errorf("%s did not match %s ", firstpoint, secondpoint)
	}

}

func TestPointSet(t *testing.T) {

	firstpoint := NewPoint(1, 1)

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

func TestJustXIndex(t *testing.T) {

	firstpoint := JustXPoint(4)
	index := JustXPoint(4)

	if !index.Match(firstpoint) {
		t.Errorf("%s did not match %s ", index, firstpoint)
	}

	secondpoint := JustXPoint(5)

	if index.Match(secondpoint) {
		t.Errorf("%s did match when it should not %s ", index, secondpoint)
	}

}

func TestJustYIndex(t *testing.T) {

	firstpoint := JustYPoint(4)
	index := JustYPoint(4)

	if !index.Match(firstpoint) {
		t.Errorf("%s did not match %s ", index, firstpoint)
	}

	secondpoint := JustYPoint(5)

	if index.Match(secondpoint) {
		t.Errorf("%s did match when it should not %s ", index, secondpoint)
	}

}

func TestJustAnyIndex(t *testing.T) {

	xpoint := JustXPoint(4)
	ypoint := JustYPoint(5)

	index := JustYPoint(4)
	index2 := JustXPoint(4)

	if index.Match(xpoint) {
		t.Errorf("%s did match when it should not %s ", index, xpoint)
	}

	if index2.Match(ypoint) {
		t.Errorf("%s did match when it should not %s ", index2, ypoint)
	}

}
