package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var r = bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscan(r, &N, &M)
	var G = make([][]bool, N+1)
	for i := range G {
		G[i] = make([]bool, N+1)
	}
	for i := 0; i < M; i++ {
		var n1, n2 int
		fmt.Fscan(r, &n1, &n2)
		G[n1][n2] = true
	}
	fmt.Println(Solve(G))
}

func Solve(G [][]bool) int {
	// for all children of root, try removing edges
	var solution int
	for i := 1; i < len(G); i++ {
		if G[i][1] {
			solution += Remove(G, 1, i)
		}
	}
	return solution
}

func Remove(G [][]bool, p, c int) int {
	// number of edges we can remove for all chilrens of node c
	var sum int

	// find all edges of c
	for i := 1; i < len(G); i++ {
		if G[i][c] {
			// i <-> c is an edge of node c to node i
			sum += Remove(G, c, i)
		}
	}

	var numChildrenOfC int
	for i := 0; i < len(G); i++ {
		if G[i][c] {
			numChildrenOfC++
		}
	}

	if numChildrenOfC%2 == 0 {
		// node c still has even number of edges. We cannot remove any edges
		return sum
	}
	// node c has odd number of edges. We remove the edge between p and c
	G[c][p] = false
	return sum + 1
}
