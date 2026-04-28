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
