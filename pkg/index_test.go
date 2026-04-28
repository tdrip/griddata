package grid

import (
	"testing"
)

func TestIndex(t *testing.T) {

	firstpoint := NewPoint(1, 1)

	index := NewIndex(firstpoint)

	if !index.GetLocation().Match(firstpoint) {
		t.Errorf("%s did not match %s ", index.GetLocation(), firstpoint)
	}

}

func TestJustXIndex(t *testing.T) {

	firstpoint := JustXPoint(4)
	index := JustXIndex(4)

	if !index.Matches(firstpoint) {
		t.Errorf("%s did not match %s ", index.GetLocation(), firstpoint)
	}

	secondpoint := JustXPoint(5)

	if index.Matches(secondpoint) {
		t.Errorf("%s did match when it should not %s ", index.GetLocation(), secondpoint)
	}

}

func TestJustYIndex(t *testing.T) {

	firstpoint := JustYPoint(4)
	index := JustYIndex(4)

	if !index.Matches(firstpoint) {
		t.Errorf("%s did not match %s ", index.GetLocation(), firstpoint)
	}

	secondpoint := JustYPoint(5)

	if index.Matches(secondpoint) {
		t.Errorf("%s did match when it should not %s ", index.GetLocation(), secondpoint)
	}

}

func TestJustAnyIndex(t *testing.T) {

	xpoint := JustXPoint(4)
	ypoint := JustYPoint(5)

	index := JustYIndex(4)
	index2 := JustXIndex(4)

	if index.Matches(xpoint) {
		t.Errorf("%s did match when it should not %s ", index.GetLocation(), xpoint)
	}

	if index2.Matches(ypoint) {
		t.Errorf("%s did match when it should not %s ", index.GetLocation(), ypoint)
	}

}
