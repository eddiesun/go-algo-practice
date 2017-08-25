package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//******************************************************************************
// DID NOT SOLVE THIS PROBLEM
//******************************************************************************

func main() {
	r := bufio.NewReader(os.Stdin)
	var board = make([][]byte, 10)
	for i := range board {
		board[i] = make([]byte, 10)
	}

	for i := range board {
		fmt.Fscan(r, &board[i])
	}
	var tmp string
	fmt.Fscan(r, &tmp)
	var words = strings.Split(tmp, ";")

	var solution = Solve(board, words)
	for i := range solution {
		fmt.Println(string(solution[i]))
	}
}

type Board [][]byte

func (b Board) Clone() Board {
	var newb = make([][]byte, 10)
	for i := range b {
		newb[i] = make([]byte, 10)
		for j := range b[i] {
			newb[i][j] = b[i][j]
		}
	}
	return newb
}

func NewBoard(b [][]byte) Board {
	return Board(b)
}

func (b Board) EmptyAt(i, j int) bool {
	return b[i][j] == '-'
}

type Slot struct {
	si       int
	sj       int
	ei       int
	ej       int
	vertical bool
}

func SlotStart(i, j int, vertical bool) Slot {
	return Slot{
		si:       i,
		sj:       j,
		vertical: vertical,
	}
}
func (s *Slot) SlotEnd(i, j int) {
	s.ei = i
	s.ej = j
}
func (s Slot) Length() int {
	if s.si > s.ei || s.sj > s.ej {
		return -1
	}
	if s.vertical {
		return s.ej - s.sj + 1
	}
	return s.ei - s.si + 1
}

func Solve(board [][]byte, words []string) [][]byte {
	var slots = buildSlots(board)
	// var slot = findSlot(board)

	fmt.Println(slots)

	return board
}

func buildSlots(board [][]byte) []Slot {
	var slots = make([]Slot, 0)
	// find horizontal slots
	for i := 0; i < 10; i++ {
		var j = 0
		for j < 10 {
			if board[i][j] == '+' {
				j++
			} else {
				// found a cell
				var s = SlotStart(i, j, false)
				for k := j + 1; k <= 10; k++ {
					if k == 10 || board[i][k] == '+' {
						(&s).SlotEnd(i, k-1)
						if s.Length() > 1 {
							slots = append(slots, s)
						}
						j = k
						break
					}
				}
			}
		}
	}
	// find vertical slots
	for j := 0; j < 10; j++ {
		var i = 0
		for i < 10 {
			if board[i][j] == '+' {
				i++
			} else {
				// found a cell
				var s = SlotStart(i, j, true)
				for k := i + 1; k <= 10; k++ {
					if k == 10 || board[k][j] == '+' {
						(&s).SlotEnd(k-1, j)
						if s.Length() > 1 {
							slots = append(slots, s)
						}
						i = k
						break
					}
				}
			}
		}
	}

	return slots
}
