package main

import (
	"fmt"
	"sync"
)

func main() {
	mtx := sync.Mutex{}
	mtx.Lock()
	c := sync.NewCond(&mtx)
	go func() {
		c.L.Lock()
		defer c.L.Unlock()
		fmt.Printf("goroutine: unlocking locked by main\n")
		c.Broadcast()
		fmt.Printf("goroutine: exiting async routine\n")
	}()
	fmt.Printf("main obtains lock, waiting for goroutine\n")
	c.Wait()
	c.L.Unlock()
	fmt.Printf("exiting\n")
}
