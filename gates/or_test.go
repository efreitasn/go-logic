package gates

import (
	"testing"

	"github.com/efreitasn/go-logic/internal/utils"
)

func TestOr(t *testing.T) {
	truthTable := []struct {
		a   bool
		b   bool
		out bool
	}{
		{true, true, true},
		{true, false, true},
		{false, false, false},
		{false, true, true},
	}

	for _, test := range truthTable {
		t.Run(utils.BoolsToStr(test.a, test.b), func(t *testing.T) {
			r := Or(test.a, test.b)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}
