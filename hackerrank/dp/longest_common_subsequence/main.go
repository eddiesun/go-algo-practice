package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	var a = make([]int, n)
	var b = make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &b[i])
	}
	var solution = Solve(a, b)
	for i, num := range solution {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
}

func Solve(a, b []int) []int {
	var cache = make([][]int, len(a)+1)
	for i := range cache {
		cache[i] = make([]int, len(b)+1)
	}

	var parents = make([][][]int, len(a)+1)
	for i := range parents {
		parents[i] = make([][]int, len(b)+1)
		for j := range parents[i] {
			parents[i][j] = make([]int, 0)
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				cache[i+1][j+1] = cache[i][j] + 1
				parents[i+1][j+1] = append(parents[i][j], a[i])
			} else {
				if cache[i][j+1] > cache[i+1][j] {
					cache[i+1][j+1] = cache[i][j+1]
					parents[i+1][j+1] = parents[i][j+1]
				} else {
					parents[i+1][j+1] = parents[i+1][j]
				}
			}
		}
	}

	var max int
	var maxseq []int
	for i := 1; i < len(cache); i++ {
		for j := 1; j < len(cache[i]); j++ {
			if cache[i][j] > max {
				max = cache[i][j]
				maxseq = parents[i][j]
			}
		}
	}
	return maxseq
}
