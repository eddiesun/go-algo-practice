package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		a   []int
		out int
	}{
		{[]int{1, 2, 100}, 197},
		{[]int{5, 3, 2}, 0},
	}

	for _, c := range cases {
		if out := Solve(c.a); out != c.out {
			t.Errorf("Solve(%v)=%v, expected %v", c.a, out, c.out)
		}
	}
}
