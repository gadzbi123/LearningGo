package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Contexts struct {
	ctx []context.Context
}

func process(ctx context.Context, wg *sync.WaitGroup, count int) {
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Println("Canceled but keep going")
			} else {
				fmt.Println("Timeout")
				goto end
			}
		default:
		}
		fmt.Println("Parsing req", i)
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Request", i, "parsed")
	}
end:
	wg.Done()
}
func interrupt(wg *sync.WaitGroup, f context.CancelFunc) {
	time.Sleep(1 * time.Second)
	f()
	wg.Done()
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go process(ctx, &wg, 10)
	wg.Add(1)
	go interrupt(&wg, cancel)
	wg.Wait()
}
