package main

import "fmt"

type Direction int

const (
	right Direction = iota + 1
	down
	left
	up
)

type Board struct {
	dir             Direction
	x, y, n, offset int
	array           [][]int
}

func newBoard(n int) *Board {
	if n < 1 {
		panic("Size can't be less then 1")
	}
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	return &Board{dir: right, n: n, array: arr}
}

func (b *Board) updateLocation() {
	switch b.dir {
	case right:
		if b.x+1 == b.n-b.offset {
			b.y += 1
			b.dir = down
		} else {
			b.x += 1
		}
	case down:
		if b.y+1 == b.n-b.offset {
			b.x -= 1
			b.dir = left
		} else {
			b.y += 1
		}
	case left:
		if b.x-b.offset == 0 {
			b.y -= 1
			b.dir = up
		} else {
			b.x -= 1
		}
	case up:
		if b.y-b.offset-1 == 0 {
			b.x += 1
			b.offset += 1
			b.dir = right
		} else {
			b.y -= 1
		}
	default:
		panic("Invalid case")
	}
}
func (b *Board) makeSpiral() {
	num := 1
	for {
		for {
			// fmt.Println(b.dir, b.x, b.y, num)
			b.array[b.y][b.x] = num
			if num == b.n*b.n {
				return
			}
			num++
			b.updateLocation()
		}
	}
}
func (b *Board) print() {
	for i := 0; i < len(b.array); i++ {
		for j := 0; j < len(b.array[0]); j++ {
			fmt.Printf("%03v ", b.array[i][j])
		}
		fmt.Println()
	}
}
func main() {
	b := newBoard(20)
	b.makeSpiral()
	b.print()
}
