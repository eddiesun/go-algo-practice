package main

import "testing"

func TestSolve(t *testing.T) {
	var cases = []struct {
		n   int
		arr []int
		out int64
	}{
		{4, []int{1, 2, 3}, 4},
		{10, []int{2, 5, 3, 6}, 5},
	}

	for _, c := range cases {
		if out := Solve(c.n, c.arr); out != c.out {
			t.Errorf("Solve(%v)=%v, expected %v", c.arr, out, c.out)
		}
	}
}
