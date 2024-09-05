package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 2, 1, 5, 3}
	for v := range slices.Chunk(nums, 2) {
		fmt.Println(v)
	}
}
