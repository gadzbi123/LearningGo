package main

import (
	"sync"
)

type safeArray struct {
	arr []int
	mtx sync.Mutex
}

func main() {
  var wg sync.WaitGroup
	var size int = 10e6
	safeArr := safeArray{arr: make([]int, size)}
	for i := 1; i <= size; i++ {
		wg.Add(1)
		go func() {
			safeArr.mtx.Lock()
			safeArr.arr = append(safeArr.arr, i)
			safeArr.mtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	// fmt.Println()
}
