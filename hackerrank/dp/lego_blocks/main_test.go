package main

import "testing"

func TestSolve(t *testing.T) {
	var cases = []struct {
		n   int
		m   int
		out int64
	}{
		{2, 2, 3},
		{3, 2, 7},
		{2, 3, 9},
		{2, 4, 27},
		{4, 4, 3375},
	}
	for _, c := range cases {
		if out := Solve(c.n, c.m); out != c.out {
			t.Errorf("Solve(%v, %v)=%v, expected %v", c.n, c.m, out, c.out)
		}
	}
}
