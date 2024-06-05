package main

import (
	"fmt"
	"math"
	"time"
)

func CopyLoop() {
	array := []int{1, 2, 3, 4, 5}
	for i, v := range array {
		array[i] = int(math.Pow(float64(v), 3))
	}
}
func SliceLoop() {
	array := []int{1, 2, 3, 4, 5}
	for _, v := range array[:] {
		v = int(math.Pow(float64(v), 3))
	}
}
func KnownLenPtrLoop() {
	array := [5]int{1, 2, 3, 4, 5}
	for _, v := range &array {
		v = int(math.Pow(float64(v), 3))
	}
}

func UnknownLenPtrLoop() {
	array := []int{1, 2, 3, 4, 5}
	for _, v := range *(&array) {
		v = int(math.Pow(float64(v), 3))
	}
}
func main() {
	t1 := time.Now()
	CopyLoop() // Slower 3x, 4x
	r1 := time.Since(t1)

	t2 := time.Now()
	SliceLoop() // Fast, but sometimes slightly slower
	r2 := time.Since(t2)

	t3 := time.Now()
	KnownLenPtrLoop() // Fast, but must known len
	r3 := time.Since(t3)

	t4 := time.Now()
	UnknownLenPtrLoop() // Fast, but unsafe and ugly
	r4 := time.Since(t4)

	fmt.Println("Time in Nanoseconds:")
	fmt.Printf(
		"CopyLoop: %v\nSliceLoop: %v\nKnownLenPtrLoop: %v\nUnknownLenPtrLoop: %v\n",
		r1.Nanoseconds(),
		r2.Nanoseconds(),
		r3.Nanoseconds(),
		r4.Nanoseconds(),
	)
}
