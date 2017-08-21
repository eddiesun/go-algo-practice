package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var t1, t2, n int64
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &t1)
	fmt.Fscan(r, &t2)
	fmt.Fscan(r, &n)
	solution := Solve(t1, t2, n)
	fmt.Println(solution.String())
}

func Solve(t1, t2, n int64) *big.Int {
	var n1 = big.NewInt(t1)
	var n2 = big.NewInt(t2)
	var fibi = big.NewInt(0)

	if n == 1 {
		return n1
	}
	if n == 2 {
		return n2
	}

	var i int64
	for i = 3; i <= n; i++ {
		fibi.Exp(n2, big.NewInt(2), nil)
		fibi.Add(fibi, n1)
		*n1 = *n2
		*n2 = *fibi
	}
	return fibi
}
