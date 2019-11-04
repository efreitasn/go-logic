package gates

import (
	"testing"

	"github.com/efreitasn/go-logic/internal/utils"
)

func TestAnd(t *testing.T) {
	truthTable := []struct {
		a   bool
		b   bool
		out bool
	}{
		{true, true, true},
		{true, false, false},
		{false, false, false},
		{false, true, false},
	}

	for _, e := range truthTable {
		t.Run(utils.BoolsToStr(e.a, e.b), func(t *testing.T) {
			r := And(e.a, e.b)

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
