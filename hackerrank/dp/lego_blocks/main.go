package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

//******************************************************************************
// DID NOT SOLVE THIS PROBLEM
//******************************************************************************

var MOD = big.NewInt(10e9 + 7)

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
	var fib = Fib(m)
	var cache = make(map[Pair]int64)
	return solve(n, m, fib, cache)
}

func solve(n, m int, fib []int64, cache map[Pair]int64) int64 {
	if m == 1 {
		return 1
	}
	if value, exists := cache[Pair{n, m}]; exists {
		return value
	}

	var firstPart big.Int
	// firstPart = fib(m)^n
	firstPart.Exp(big.NewInt(fib[m]), big.NewInt(int64(n)), MOD)
	var secondPart big.Int
	for i := 1; i <= m-1; i++ {
		var tmp1, tmp2, tmp3 big.Int
		// tmp1 = fib(m-i)^n
		tmp1.Exp(big.NewInt(fib[m-i]), big.NewInt(int64(n)), MOD)
		// tmp2 = solve(n,m-i)
		tmp2.SetInt64(solve(n, i, fib, cache))
		// tmp3 = tmp1 * tmp 2 = fib(m-i)^n * solve(n,m-i)
		tmp3.Mul(&tmp1, &tmp2).Mod(&tmp3, MOD)
		// secondPart = secondPart + tmp3
		secondPart.Add(&secondPart, &tmp3).Mod(&secondPart, MOD)
	}
	var result big.Int
	result.Sub(&firstPart, &secondPart)

	cache[Pair{n, m}] = result.Int64()
	return result.Int64()
}

func Fib(m int) []int64 {
	var f = make([]int64, max(4, m+1))
	f[0] = 1
	f[1] = 1
	f[2] = 2
	f[3] = 4
	for i := 4; i <= m; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3] + f[i-4]
	}
	return f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
