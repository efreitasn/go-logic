package nand

import (
	"testing"

	"github.com/efreitasn/go-logic/internal/utils"
)

func TestExec(t *testing.T) {
	truthTable := []struct {
		a   bool
		b   bool
		out bool
	}{
		{true, true, false},
		{true, false, true},
		{false, false, true},
		{false, true, true},
	}

	for _, e := range truthTable {
		t.Run(utils.BoolsToStr(e.a, e.b), func(t *testing.T) {
			r := Exec(e.a, e.b)

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
