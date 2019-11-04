package gates

import (
	"testing"

	"github.com/efreitasn/go-logic/internal/utils"
)

func TestNot(t *testing.T) {
	var truthTable = []struct {
		in  bool
		out bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range truthTable {
		t.Run(utils.BoolsToStr(test.in), func(t *testing.T) {
			r := Not(test.in)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}
