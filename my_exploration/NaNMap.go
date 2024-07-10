package main

import (
	"fmt"
	"math"
	"math/rand"

	"golang.org/x/exp/maps"
)

func main() {
	nan := math.NaN()
	my_map := map[float64]int{}
	for i := 0; i < 10; i++ {
		my_map[nan] = i
	}
	// This method is less random like [2 3 4 6 0 1 5 7 8 9]
	fmt.Println(maps.Values(my_map)) // It puts values while has space, then finds different place

	s := rand.Perm(10)
	fmt.Println(s) // This output is super random like [9 2 6 3 0 7 8 1 4 5]
}
