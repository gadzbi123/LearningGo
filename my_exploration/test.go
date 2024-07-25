package main

import (
	"fmt"
	"time"
)

func main() {
	x := func(v int) {
		fmt.Println("XD", v)
	}
	go x(1)
	go x(2)
	go x(3)
	time.Sleep(1 * time.Second)
}
