package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var T int
	r := bufio.NewReader(os.Stdin)
	for fmt.Fscan(r, &T); T > 0; T-- {
		var a, b string
		fmt.Fscan(r, &a, &b)
		output := Solve(a, b)
		if output {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func Solve(a, b string) bool {
	var cache = make([][]bool, len(a)+1)
	for i := range cache {
		cache[i] = make([]bool, len(b)+1)
	}

	for i := len(cache) - 1; i >= 0; i-- {
		if i == len(cache)-1 {
			cache[i][len(b)] = true
			continue
		}
		// for last letter in a, if uppercase, you cannot discard it to make it true
		// if lowercase, you can discard it to make it true
		if unicode.IsUpper(rune(a[i])) {
			break
		} else {
			cache[i][len(b)] = true
		}
	}

	var i, j int

	for i = len(a) - 1; i >= 0; i-- {
		for j = len(b) - 1; j >= 0; j-- {
			if unicode.IsUpper(rune(a[i])) {
				if a[i] == b[j] {
					cache[i][j] = cache[i+1][j+1]
				} else {
					cache[i][j] = false
				}
			} else {
				if byte(unicode.ToUpper(rune(a[i]))) == b[j] {
					cache[i][j] = cache[i+1][j+1] || cache[i+1][j]
				} else {
					cache[i][j] = cache[i+1][j]
				}
			}
		}
	}
	// fmt.Println(cache)
	return cache[0][0]
}
