package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(3)
	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		ch <- "Aleksander"
	}()
	go func() {
		defer wg.Done()
		ch <- "Grzegorz"
	}()
	go func() {
		defer wg.Done()
		ch <- "Andrzej"
	}()
	for name := range ch {
		fmt.Println("Person", name, "is calling")
	}
	v, ok := <-ch
	fmt.Printf("%q, %v", v, ok)
}
