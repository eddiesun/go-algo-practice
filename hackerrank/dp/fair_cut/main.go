package main

import (
	"bufio"
	"fmt"
	"os"
)

//******************************************************************************
// DID NOT SOLVE THIS PROBLEM
//******************************************************************************

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(r, &n, &k)
	var a []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	solution := Solve(a, k)
	fmt.Println(solution)
}

var set []int

func Solve(s []int, k int) int {
  set = s
	var i = 0
	var p, q = 0, 0 // num ints that Li has and num ints that Lu has. Sum of p and q should be len(s). p <= k and p eventually must equal to k
	var unfairness = 0
	return solve(i, p, q, unfairness)
}

func solve(i, p, q, unfairness int) int {
  var current = set[i]
  // if give current to Li
  var unfairness1 = calcUnfairness(unfairness, current, setOfLu)
  solve(i+1, p+1, q, unfairness1)
  // if give current to Lu
  var unfairness2
}
