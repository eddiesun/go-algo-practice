package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(r, &T)

	for ; T > 0; T-- {
		var n int
		fmt.Fscan(r, &n)
		var a = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		fmt.Println(Solve(a))
	}
}

type Cell struct {
	profit   int
	maxIndex int
}

func Solve(a []int) int {
	var cache = make([]Cell, len(a)+1)
	for i := len(a) - 1; i >= 0; i-- {
		var futureProfit = cache[i+1].profit
		var futureMaxIndex = cache[i+1].maxIndex

		if i == len(a)-1 || a[i] > a[futureMaxIndex] {
			cache[i].profit = futureProfit
			cache[i].maxIndex = i
		} else {
			cache[i].profit = futureProfit + a[futureMaxIndex] - a[i]
			cache[i].maxIndex = futureMaxIndex
		}
	}
	// fmt.Println(cache)
	return cache[0].profit
}
