package main

import "testing"

func TestSolve(t *testing.T) {
	var testcases = []struct {
		t1     int64
		t2     int64
		n      int64
		output string
	}{
		{0, 1, 1, "0"},
		{0, 1, 2, "1"},
		{0, 1, 3, "1"},
		{0, 1, 4, "2"},
		{0, 1, 5, "5"},
	}

	for _, c := range testcases {
		output := Solve(c.t1, c.t2, c.n).String()
		if c.output != output {
			t.Errorf("Solve(%v, %v, %v)=%v, expected %v", c.t1, c.t2, c.n, output, c.output)
		}
	}
}
