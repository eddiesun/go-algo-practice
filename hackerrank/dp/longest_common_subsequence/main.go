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

type Cell struct {
	Len   int
	PrevI int
	PrevJ int
	hit   bool
}

func Solve(a, b []int) []int {
	var cache = make([][]Cell, len(a)+1)
	for i := range cache {
		cache[i] = make([]Cell, len(b)+1)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				cache[i+1][j+1].Len = cache[i][j].Len + 1
				cache[i+1][j+1].PrevI = i
				cache[i+1][j+1].PrevJ = j
				cache[i+1][j+1].hit = true
			} else {
				if cache[i][j+1].Len > cache[i+1][j].Len {
					cache[i+1][j+1].Len = cache[i][j+1].Len
					cache[i+1][j+1].PrevI = i
					cache[i+1][j+1].PrevJ = j + 1
				} else {
					cache[i+1][j+1].Len = cache[i+1][j].Len
					cache[i+1][j+1].PrevI = i + 1
					cache[i+1][j+1].PrevJ = j
				}
			}
		}
	}

	// construct the max sequence by following PrevI and PrevJ
	var solution []int
	var i = len(cache) - 1
	var j = len(cache[i]) - 1
	var cell = cache[i][j]
	for {
		if i == 0 || j == 0 {
			break
		}
		if cell.hit {
			solution = append(solution, a[i-1])
		}
		i, j = cell.PrevI, cell.PrevJ
		cell = cache[i][j]
	}
	// reverse the sequence to get final answer
	for i := 0; i < len(solution)/2; i++ {
		solution[i], solution[len(solution)-1-i] = solution[len(solution)-1-i], solution[i]
	}
	return solution
}
