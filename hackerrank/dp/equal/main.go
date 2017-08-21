package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	var r = bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var N int
		fmt.Fscan(r, &N)
		var a = make([]int, N)
		for i := 0; i < len(a); i++ {
			fmt.Fscan(r, &a[i])
		}
		var solution = solve(a)
		fmt.Println(solution)
	}
}

func solve(a []int) int {
	// reduce to difference array
	b0 := prep(a, 0)
	b1 := prep(a, 1)
	b2 := prep(a, 2)

	var count0 int
	for _, v := range b0 {
		count0 += reduce(v)
	}

	var count1 int
	for _, v := range b1 {
		count1 += reduce(v)
	}

	var count2 int
	for _, v := range b2 {
		count2 += reduce(v)
	}

	return min(count0, count1, count2)
}

func prep(a []int, offset int) []int {
	var b = make([]int, len(a))

	// find min and element-wise minus min
	var min = a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	for i := range b {
		b[i] = a[i] - min + offset
	}

	return b
}

func reduce(n int) int {
	var count int
	count += n / 5
	n %= 5
	count += n / 2
	n %= 2
	count += n
	return count
}

func min(a, b, c int) int {
	var min = 0
	if a < b {
		min = a
	} else {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
