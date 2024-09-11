package main

import (
	"fmt"
	"sync"
)

func main() {
	sync.OnceFunc(func() {
		fmt.Println("XD")
	})
}
