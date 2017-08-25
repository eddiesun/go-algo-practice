package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var MOD int64 = 10e9 + 7

func main() {
	var r = bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, m int
		fmt.Fscan(r, &n, &m)
		fmt.Println(Solve(n, m))
	}
}

type Pair struct {
	n int
	m int
}

func Solve(n, m int) int64 {
	var foundation = buildFoundation(m)
	var cache = make(map[Pair]int64)
	return solve(n, m, foundation, cache)
}

func solve(n, m int, foundation []int64, cache map[Pair]int64) int64 {
	if m == 1 {
		return 1
	}
	var firstPart big.Int
	firstPart.Exp(big.NewInt(foundation[m]), big.NewInt(int64(n)), big.NewInt(MOD))
	var secondPart big.Int
	for i := 1; i <= m-1; i++ {
		var tmp big.Int
		tmp.Mul(
			big.NewInt(solve(n, i, foundation, cache)),
			big.NewInt(solve(n, m-i, foundation, cache)),
		).Mod(&tmp, big.NewInt(MOD))
		secondPart.Add(&secondPart, &tmp).Mod(&secondPart, big.NewInt(MOD))
	}
	var result big.Int
	result.Sub(&firstPart, &secondPart)
	cache[Pair{n, m}] = result.Int64()
	return result.Int64()
}

func buildFoundation(m int) []int64 {
	var f = make([]int64, max(4, m+1))
	f[0] = 1
	f[1] = 1
	f[2] = 2
	f[3] = 4
	for i := 4; i <= m; i++ {
		f[i] = f[i-1] + f[i-1] + f[i-3] + f[i-4]
	}
	return f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
