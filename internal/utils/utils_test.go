package utils

import (
	"strconv"
	"testing"
)

func TestBoolsToStr(t *testing.T) {
	tests := []struct {
		in  []bool
		out string
	}{
		{[]bool{true, false, true}, "true_false_true"},
		{[]bool{false, false, true}, "false_false_true"},
		{[]bool{false, false, false}, "false_false_false"},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := BoolsToStr(test.in...)

			if r != test.out {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}
}
