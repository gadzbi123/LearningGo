package main

import (
	"fmt"
	"time"
)

const TIMEOUT_LEN = 3 //1
func main() {
	start := time.NewTimer(TIMEOUT_LEN * time.Second)
	request := time.NewTimer(2 * time.Second)
	done := make(chan bool)
	fmt.Println("start")

	go func() {
		<-request.C
		fmt.Println("req took 2 sec")
		done <- true
	}()
F:
	for {
		select {
		case <-done:
			break F
		case <-start.C:
			fmt.Printf("Req didn't come after %v secs, ABORT\n", TIMEOUT_LEN)
			break F
		default:
		}
	}
}
