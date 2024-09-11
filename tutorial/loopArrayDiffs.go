package main

import (
	"fmt"
)

func Ptr[T any](v T) *T {
	return &v
}

func main() {
	standard := []int{1, 2, 3}
	for _, v := range standard {
		v += 1
		fmt.Printf("standard: %q\n", v)
	}
	fmt.Printf("full standard: %q\n", standard)

	for _, v := range standard[:] {
		v += 1
		fmt.Printf("standard slice: %q\n", v)
	}
	fmt.Printf("full standard slice: %q\n", standard)

	slicePtr := &[]int{1, 2, 3}
	for _, v := range *slicePtr {
		v += 1
		fmt.Printf("slicePtr: %q\n", v)
	}
	fmt.Printf("full slicePtr: %q\n", slicePtr)

	arrOfPtr := []*int{Ptr(1), Ptr(2), Ptr(3)}
	for _, v := range arrOfPtr {
		*v += 1
		fmt.Printf("arrOfPtr: %q\n", *v)
	}
	fmt.Printf("first arrOfPtr: %v\n", *arrOfPtr[0])
}
