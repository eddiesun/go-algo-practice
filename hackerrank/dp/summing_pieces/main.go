package main

import (
	"bufio"
	"fmt"
	"os"
)

//******************************************************************************
// DID NOT SOLVE THIS PROBLEM
//******************************************************************************

var M int64 = 10e9 + 7

func main() {
	r := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(r, &N)
	var a = make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &a[i])
	}
	fmt.Println(Solve(a))
}

func Solve(a []int64) int64 {
	var prefixSum = make([]int64, len(a))
	for i := 0; i < len(prefixSum); i++ {
		if i == 0 {
			prefixSum[i] = a[i] % M
		} else {
			prefixSum[i] = (prefixSum[i-1] + a[i]) % M
		}
	}

	var partialResult int64
	var i, j int64
	for i = 0; i < int64(len(a)); i++ {
		if i == 0 {
			partialResult = a[i]
		} else {
			var tmp int64
			for j = i - 1; j >= -1; j-- {
				if j == -1 {
					tmp = (tmp + prefixSum[i]*(i-j)) % M
				} else {
					tmp = (tmp + (prefixSum[i]-prefixSum[j])*(i-j)) % M
				}
			}
			partialResult = (partialResult + tmp) % M
		}
	}

	return partialResult
}
