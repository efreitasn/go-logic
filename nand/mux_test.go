package nand

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/efreitasn/go-logic/internal/utils"
)

func TestMux(t *testing.T) {
	truthTable := []struct {
		a   bool
		b   bool
		s   bool
		out bool
	}{
		{true, true, true, true},
		{true, true, false, true},
		{true, false, false, true},
		{true, false, true, false},
		{false, false, false, false},
		{false, false, true, false},
		{false, true, true, true},
		{false, true, false, false},
	}

	for _, e := range truthTable {
		t.Run(utils.BoolsToStr(e.a, e.b, e.s), func(t *testing.T) {
			r := Mux(e.a, e.b, e.s)

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}

func TestMultiBitMux(t *testing.T) {
	truthTable := []struct {
		a   []bool
		b   []bool
		s   bool
		out []bool
	}{
		{
			a:   []bool{true, true, true},
			b:   []bool{false, true, false},
			s:   false,
			out: []bool{true, true, true},
		},
		{
			a:   []bool{false, true, false},
			b:   []bool{false, false, false},
			s:   true,
			out: []bool{false, false, false},
		},
	}

	for _, e := range truthTable {
		t.Run(fmt.Sprintf("s=%v", e.s), func(t *testing.T) {
			r := MultiBitMux(e.a, e.b, e.s)

			if !reflect.DeepEqual(r, e.out) {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
