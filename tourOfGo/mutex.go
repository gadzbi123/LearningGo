package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mtx   sync.Mutex
	mapty map[string]int
}

func (sc *SafeCounter) Inc(key string) {
	sc.mtx.Lock()
	sc.mapty[key]++
	sc.mtx.Unlock()
}

func (sc *SafeCounter) read() {
	// Uncomment 2 lines to not crash
	// sc.mtx.Lock()
	fmt.Println(sc.mapty)
	// sc.mtx.Unlock()
}

func main() {
	sc := &SafeCounter{mapty: make(map[string]int)}
	go func() {
		for i := 0; i < 10e6; i++ {
			sc.Inc("my_key")
		}
	}()
	t := time.Duration(5)
	for {
		sc.read()
		time.Sleep(t * time.Nanosecond)
		t *= 5
	}
}
