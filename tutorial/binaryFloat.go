package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var myF float32 = 3
	x := *(*uint32)(unsafe.Pointer(&myF))
	fBits := math.Float32bits(myF)
	fmt.Printf("%15v -> %032b\n", myF, fBits)
	x = x | 0x80000000 //== x = x | (1 << 31)
	// y := uint32(0xFFFFFFFF ^ (1 << 30))
	// yBits := strconv.FormatUint(*(*uint64)(unsafe.Pointer(&y)), 2)
	// fmt.Printf("%15v -> %32b\n", y, y)
	// x = x & (y)
	myF = *(*float32)(unsafe.Pointer(&x))
	fBits = math.Float32bits(myF)
	fmt.Printf("%15v -> %032b\n", myF, fBits)
}
