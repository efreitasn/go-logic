package adder

import (
	"fmt"
	"testing"
)

func TestFull(t *testing.T) {
	truthTable := []struct {
		a    uint8
		b    uint8
		cIn  uint8
		cOut uint8
		s    uint8
	}{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 1, 1, 1},
	}

	for _, row := range truthTable {
		t.Run(fmt.Sprintf("%v_%v_%v", row.a, row.b, row.cIn), func(t *testing.T) {
			cOut, s := Full(row.a, row.b, row.cIn)

			if cOut != row.cOut {
				t.Errorf("got %v as cOut, want %v", cOut, row.cOut)
			}

			if s != row.s {
				t.Errorf("got %v as s, want %v", s, row.s)
			}
		})
	}
}
