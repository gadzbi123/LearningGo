package main

import (
	"fmt"
	"maps"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc/1024, "KB allocations")
}

func main() {
	myMap := make(map[int][128]byte, 1000)
	printAllocs()
	for i := 0; i < 1000; i++ {
		myMap[i] = [128]byte{}
		arr := myMap[i]
		copy(arr[:], []byte("Hello"))
	}
	printAllocs()
	for i := 0; i < 1000; i += 2 {
		delete(myMap, i)
	}
	printAllocs()
	myMap2 := map[int][128]byte{}
	maps.Copy(myMap2, myMap)
	runtime.GC()
	printAllocs()
	fmt.Println("map2 len:", len(myMap2))
	runtime.KeepAlive(myMap2)
}
