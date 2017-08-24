package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		a   []int
		b   []int
		out []int
	}{
		{[]int{1, 2, 3, 4, 1}, []int{3, 4, 1, 2, 1, 3}, []int{3, 4, 1}},
		// {[]int{1, 2}, 9},
	}

	for _, c := range cases {
		if out := Solve(c.a, c.b); !reflect.DeepEqual(out, c.out) {
			t.Errorf("Solve(%v, %v)=%v, expected %v", c.a, c.b, out, c.out)
		}
	}
}
