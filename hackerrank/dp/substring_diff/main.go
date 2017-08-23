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
	var N int
	fmt.Fscan(r, &N)
	for ; N > 0; N-- {
		var k int
		var a, b string
		fmt.Fscan(r, &k, &a, &b)
		var solution = Solve(k, a, b)
		fmt.Println(solution)
	}
}

type Cell struct {
	K   int
	Len int
}

func Solve(k int, a, b string) int {
	var cache = make([][][]Cell, len(a)+1)
	for i := range cache {
		cache[i] = make([][]Cell, len(b)+1)
	}
	for _, column := range cache {
		for j := range column {
			column[j] = make([]Cell, 0)
		}
	}
	for i := 0; i < len(cache); i++ {
		cache[i][0] = append(cache[i][0], Cell{K: k, Len: 0})
	}
	for j := 1; j < len(cache[0]); j++ {
		cache[0][j] = append(cache[0][j], Cell{K: k, Len: 0})
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {

			if a[i] == b[j] {
				for p := range cache[i][j] {
					cache[i+1][j+1] = append(cache[i+1][j+1], Cell{
						K:   cache[i][j][p].K,
						Len: cache[i][j][p].Len + 1,
					})
				}
			} else {
				cache[i+1][j+1] = append(cache[i+1][j+1], Cell{
					K:   k,
					Len: 0,
				})
				for p := range cache[i][j] {
					if cache[i][j][p].K > 0 {
						cache[i+1][j+1] = append(cache[i+1][j+1], Cell{
							K:   cache[i][j][p].K - 1,
							Len: cache[i][j][p].Len + 1,
						})
					}
				}
			}
		}
	}

	// now we scan the entire 2D array and find the cell that has max Len field
	var max_i, max_j, max int
	for i := 1; i < len(cache); i++ {
		for j := 1; j < len(cache[i]); j++ {
			for p := range cache[i][j] {
				if cache[i][j][p].Len > max {
					max = cache[i][j][p].Len
					max_i = i - 1
					max_j = j - 1
				}
			}
		}
	}
	_ = max_i
	_ = max_j
	// fmt.Println(max_i, max_j)
	// fmt.Println(cache)
	return max
}
