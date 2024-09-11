package main

import (
	"fmt"
	"time"
)

func main() {
	first := time.Tick(1 * time.Second)
	last := time.After(5 * time.Second)
	start := true
	cur := 1
	for start {
		select {
		case <-first:
			fmt.Println(cur, "sec passed")
			cur++
		case <-last:
			start = false
		}
	}
}
