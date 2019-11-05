package adder

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a   int
		b   int
		out int
	}{
		{20, 30, 50},
		{1000, 2, 1002},
		{0, 1, 1},
		{1, 1, 2},
		{0, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v+%v", test.a, test.b), func(t *testing.T) {
			r := Add(test.a, test.b)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}

func TestAddFullAdder(t *testing.T) {
	tests := []struct {
		a   int
		b   int
		out int
	}{
		{20, 30, 50},
		{1000, 2, 1002},
		{0, 1, 1},
		{1, 1, 2},
		{0, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v+%v", test.a, test.b), func(t *testing.T) {
			r := AddFullAdder(test.a, test.b)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}
