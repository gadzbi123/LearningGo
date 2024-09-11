package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(10))
	}
	values := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
	fmt.Println(values)
	slices.SortStableFunc(values, func(a, b int) int {
		if a <= b {
			return 1
		}
		return -1
	})
	fmt.Println(values)
	if !sort.IntsAreSorted(values) {
		slices.Sort(values)
		fmt.Println(values)
	}
}
