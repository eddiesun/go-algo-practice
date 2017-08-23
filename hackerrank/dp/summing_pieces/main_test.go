package main

import "testing"

func TestSolve(t *testing.T) {
	var cases = []struct {
		arr []int64
		out int64
	}{
		{[]int64{1}, 1},
		{[]int64{1, 2}, 9},
		{[]int64{1, 2, 3}, 40},
		{[]int64{4}, 4},
		{[]int64{4, 2}, 18},
		{[]int64{4, 2, 9}, 94},
		{[]int64{4, 2, 9, 10}, 305},
		{[]int64{4, 2, 9, 10, 1}, 971},
	}

	for _, c := range cases {
		if out := Solve(c.arr); out != c.out {
			t.Errorf("Solve(%v)=%v, expected %v", c.arr, out, c.out)
		}
	}
}
