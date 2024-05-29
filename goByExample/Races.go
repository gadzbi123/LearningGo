package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var size int = 10e6
	arr := make([]int, 0, size)
	for i := 1; i <= size; i++ {
		wg.Add(1)
		go func() {
			arr = append(arr, i)
			wg.Done()
		}()
	}
	wg.Wait()
	// fmt.Println()
}
