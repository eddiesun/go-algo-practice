package main

import "testing"

func TestSolve(t *testing.T) {
	var cases = []struct {
		x   []int
		k   int
		out int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 2},
		{[]int{7, 2, 4, 6, 5, 9, 12, 11}, 2, 3},
	}
	for _, c := range cases {
		if out := Solve(c.x, c.k); out != c.out {
			t.Errorf("Solve(%v, %v)=%v, expected %v", c.x, c.x, out, c.out)
		}
	}
}
