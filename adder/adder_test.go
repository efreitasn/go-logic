package adder

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	truthTable := []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{10, 20, 30},
		{50, 50, 100},
		{223, 876, 1099},
	}

	for _, e := range truthTable {
		t.Run(fmt.Sprintf("%v + %v", e.a, e.b), func(t *testing.T) {
			res := Add(e.a, e.b)

			if res != e.out {
				t.Errorf("got %v, want %v", res, e.out)
			}
		})
	}
}
