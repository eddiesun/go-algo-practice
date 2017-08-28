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
	var r = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	var nodes = make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{
			value: i + 1,
		}
	}
	for i := 0; i < n; i++ {
		var left, right int
		fmt.Fscan(r, &left, &right)
		if left >= 0 {
			nodes[i].left = nodes[left-1]
		}
		if right >= 0 {
			nodes[i].right = nodes[right-1]
		}
	}
	var numk int
	fmt.Fscan(r, &numk)
	for i := 0; i < numk; i++ {
		var k int
		fmt.Fscan(r, &k)
		Solve(nodes[0], k)
	}
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solve(root *Node, k int) {
	for i := 1; i <= k; i++ {
		Swap(root, i*k)
	}
	Print(root)
	fmt.Println()
}

func Swap(node *Node, height int) {
	if node == nil {
		return
	}
	if height == 1 {
		node.left, node.right = node.right, node.left
	} else {
		Swap(node.left, height-1)
		Swap(node.right, height-1)
	}
}

func Print(node *Node) {
	if node == nil {
		return
	}
	Print(node.left)
	fmt.Printf("%v ", node.value)
	Print(node.right)
}
