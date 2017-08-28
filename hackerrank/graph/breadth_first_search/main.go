package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//******************************************************************************
// DID NOT SOLVE THIS PROBLEM
//******************************************************************************

func main() {
	var r = bufio.NewReader(os.Stdin)
	var q int
	fmt.Fscan(r, &q)
	for i := 0; i < q; i++ {
		var n, m int
		fmt.Fscan(r, &n, &m)
		var g = make([][]int, n)
		for k := range g {
			g[k] = make([]int, n)
			for p := range g[k] {
				g[k][p] = math.MaxInt32
			}
			g[k][k] = 0
		}
		for j := 0; j < m; j++ {
			var from, to int
			fmt.Fscan(r, &from, &to)
			g[from-1][to-1] = 6
			g[to-1][from-1] = 6
		}
		var start int
		fmt.Fscan(r, &start)
		var solution = Solve(g, start-1)
		for i, s := range solution {
			if i == 0 {
				fmt.Printf("%v", s)
			} else {
				fmt.Printf(" %v", s)
			}
		}
		fmt.Println()
	}
}

func Solve(g [][]int, start int) []int {
	var queue = NewQueue(len(g))
	var visited = make([]bool, len(g))
	var parents = make([]int, len(g))
	parents[start] = start
	var dist = make([]int, len(g))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	queue.Enqueue(start)
	for !queue.IsEmpty() {
		var node = queue.Dequeue()
		if !visited[node] {
			visited[node] = true
			dist[node] = dist[parents[node]] + g[parents[node]][node]
			for i := range g {
				if i != node && g[i][node] < math.MaxInt32 {
					parents[i] = node
					queue.Enqueue(i)
				}
			}
		}
	}

	for i := range dist {
		if dist[i] == math.MaxInt32 {
			dist[i] = -1
		}
	}

	return append(dist[:start], dist[start+1:]...)
}

type Queue struct {
	queue []int
	head  int
	tail  int
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) Enqueue(node int) {
	q.queue[q.tail] = node
	q.tail = (q.tail + 1) % len(q.queue)
}

func (q *Queue) Dequeue() int {
	var node = q.queue[q.head]
	q.head = (q.head + 1) % len(q.queue)
	return node
}

func NewQueue(n int) *Queue {
	var q = new(Queue)
	q.queue = make([]int, n)
	q.head = 0
	q.tail = 0
	return q
}
