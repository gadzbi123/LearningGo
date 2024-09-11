package main

import (
	"fmt"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc/1024, "KB allocations")
}

type Foo struct {
	v []byte
}

func take2(foos []Foo) []Foo {
	foos2 := make([]Foo, 2)
	copy(foos2, foos)
	return foos2
}
func main() {
	sliceFoo := make([]Foo, 1000)
	printAllocs()
	for i := 0; i < 1000; i++ {
		sliceFoo[i].v = make([]byte, 1024*1024)
	}
	printAllocs()
	sliceFoo2 := take2(sliceFoo)
	runtime.GC()
	printAllocs()
	runtime.KeepAlive(sliceFoo2)
}
