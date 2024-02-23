package main

import (
	"fmt"
	"math/rand"
)

const TREE_SIZE = 10

type Tree struct {
	left, right *Tree
	value       int
}

func (t *Tree) new(n int) *Tree {
	for _, x := range rand.Perm(TREE_SIZE) {
		t = insert(t, (n)*x)
	}
	// fmt.Println(t.String())
	return t
}

func insert(t *Tree, value int) *Tree {
	if t == nil {
		t = &Tree{nil, nil, value}
		return t
	}
	if value <= t.value {
		t.left = insert(t.left, value)
	} else {
		t.right = insert(t.right, value)
	}
	return t
}

func (t *Tree) String() string {
	var s string
	if t == nil {
		return ""
	}
	s += fmt.Sprintf("%v %v %v", t.left.String(), fmt.Sprint(t.value), t.right.String())
	return s
}

func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < TREE_SIZE; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func Walk(t *Tree, out chan int) {
	if t == nil {
		return
	}
	Walk(t.left, out)
	out <- t.value
	Walk(t.right, out)
}

func main() {
	var t1, t2 *Tree
	t1 = t1.new(1)
	t2 = t2.new(1)
	fmt.Println(t1.String())
	fmt.Println(t2.String())
	isSame := Same(t1, t2)
	fmt.Println(isSame)
}
