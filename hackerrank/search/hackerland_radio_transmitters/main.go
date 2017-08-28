package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var r = bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(r, &n, &k)
	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x[i])
	}
	var solution = Solve(x, k)
	fmt.Println(solution)
}

func Solve(x []int, k int) int {
	sort.Ints(x)
	var count int
	for i := 0; i < len(x); {
		var j = findNextCityThatCanCoverMe(x, i, k)
		// fmt.Printf("x[%v]=%v can cover x[%v]=%v\n", j, x[j], i, x[i])
		count++
		var p = findNextUncoveredCity(x, j, k)
		// all cities are covered. No uncovered city
		if p < 0 {
			break
		}
		// fmt.Printf("next uncovered city after x[%v]=%v is x[%v]=%v\n", i, x[i], p, x[p])
		i = p
	}
	return count
}

func findNextCityThatCanCoverMe(x []int, i, k int) int {
	for q := i + 1; q < len(x); q++ {
		if x[q]-k > x[i] {
			return q - 1
		}
	}
	return len(x) - 1 // last city
}

func findNextUncoveredCity(x []int, i, k int) int {
	for q := i + 1; q < len(x); q++ {
		if x[q] > x[i]+k {
			return q
		}
	}
	return -1
}
