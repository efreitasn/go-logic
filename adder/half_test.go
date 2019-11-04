package adder

import (
	"fmt"
	"testing"
)

func TestHalf(t *testing.T) {
	truthTable := []struct {
		a uint8
		b uint8
		c uint8
		s uint8
	}{
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 0},
	}

	for _, row := range truthTable {
		t.Run(fmt.Sprintf("%v_%v", row.a, row.b), func(t *testing.T) {
			c, s := Half(row.a, row.b)

			if c != row.c {
				t.Errorf("got %v as c, want %v", c, row.c)
			}

			if s != row.s {
				t.Errorf("got %v as s, want %v", s, row.s)
			}
		})
	}
}
