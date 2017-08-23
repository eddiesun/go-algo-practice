package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, i int
	fmt.Fscan(r, &n, &m)
	var coins = make([]int, m)
	for i = 0; i < m; i++ {
		fmt.Fscan(r, &coins[i])
	}
	fmt.Println(Solve(n, coins))
}

func Solve(n int, coins []int) int64 {
	// initialize a cache
	var cache = make([][]int64, n+1)
	for i := range cache {
		cache[i] = make([]int64, len(coins)+1)
		if i == 0 {
			for j := range cache[i] {
				cache[0][j] = 1
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j < len(cache[i]); j++ {
			var coinValue = coins[j-1]
			var sum int64
			for k := 0; i-k*coinValue >= 0; k++ {
				sum += cache[i-k*coinValue][j-1]
			}
			cache[i][j] = sum
		}
	}

	// fmt.Println(cache)

	return cache[n][len(cache[n])-1]
}
