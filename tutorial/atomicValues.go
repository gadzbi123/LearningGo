package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

func rec(v int) {
	if v == 0 {
		return
	}
	vv := (uintptr)(unsafe.Pointer(&v))
	fmt.Printf("routine %v: %#x\n", v, vv) // difference between stack is 0x6000
	// which is 24_576 bytes of size of stack for goroutine
	topstack.Store(vv)
	go rec(v - 1)
}

var topstack = atomic.Uintptr{}

func main() {
	rec(11)
	time.Sleep(1)
	fmt.Printf("%x", topstack.Load())
}
