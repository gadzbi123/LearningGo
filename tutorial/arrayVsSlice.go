package main

import "fmt"

type Cargo struct {
	key   int
	boxes []int
}

func main() {
	arrayCargo := [2]Cargo{
		Cargo{key: 1, boxes: []int{1, 2, 3}},
		Cargo{key: 2, boxes: []int{4, 5, 6}},
	}
	// can't append to array, only to slice
	// arrayCargo = append(arrayCargo[:],
	// 	Cargo{key: 2, boxes: []int{4, 5, 6}},
	// )
	sliceCargo := []Cargo{
		Cargo{key: 1, boxes: []int{1, 2, 3}},
		Cargo{key: 2, boxes: []int{4, 5, 6}},
	}
	sliceCargo = append(sliceCargo[:],
		Cargo{key: 2, boxes: []int{4, 5, 6}},
	)
	fmt.Println(arrayCargo)
	fmt.Println(sliceCargo)
}
