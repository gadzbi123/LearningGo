package main

import (
	"fmt"
	"sync"
	"time"
)

type counterMtx struct {
	counter *int
	mtx     sync.RWMutex
}

func main() {
	val := 0
	my_counter := &counterMtx{counter: &val}
	for _ = range time.Tick(time.Nanosecond) {
		for i := 0; i < 5; i++ {
			go readConcurently(my_counter, i)
		}
		go writeConcurently(my_counter)
	}
}

func readConcurently(ctr *counterMtx, i int) {
	ctr.mtx.RLock()
	fmt.Printf("Worker %v reading: %v\n", i, *ctr.counter)
	ctr.mtx.RUnlock()
}

func writeConcurently(ctr *counterMtx) {
	// ctr.mtx.Lock()
	fmt.Println("Writing ctr")
	*ctr.counter += 1
	// ctr.mtx.Unlock()
}
