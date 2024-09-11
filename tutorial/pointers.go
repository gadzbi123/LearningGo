package main

import (
	"fmt"
)

type Robot struct {
	id   uint8
	data []uint8
}

func inc(x *int) {
	fmt.Printf("x: %p\n", x)
	*x++
}
func main() {
	first := []int{5, 3, 2}
	temp := &first
	var second **[]int = &temp
	fmt.Printf("%p,%p,%p\n", first, *temp, **second)
	first = append(first, 0)
	// first = slices.Concat(first, first[1:2])
	fmt.Println(first, *temp, **second)
	fmt.Printf("%p,%p,%p\n", &first[1], *temp, **second)

	var borrowed int
	fmt.Printf("borrowed: %p\n", &borrowed)
	inc(&borrowed)
	fmt.Println(borrowed)
}
