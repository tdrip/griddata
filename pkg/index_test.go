package grid

import (
	"testing"
)

func TestIndex(t *testing.T) {

	firstpoint := CreatePoint(1, 1)

	index := CreateIndex(firstpoint)

	if !index.GetPosition().Match(firstpoint) {
		t.Errorf("%s did not match %s ", index.GetPosition(), firstpoint)
	}

}
