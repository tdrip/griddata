package grid

import (
	"testing"
)

func TestPoint(t *testing.T) {

	firstpoint := CreatePoint(1, 1)

	if firstpoint.GetX() != 1 || firstpoint.GetY() != 1 {
		t.Errorf("X was %d not 1 - Y was %d not 1", firstpoint.GetX(), firstpoint.GetY())
	}

	secondpoint := CreatePoint(1, 1)

	if !firstpoint.Match(secondpoint) {
		t.Errorf("%s did not match %s ", firstpoint, secondpoint)
	}

}
