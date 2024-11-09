package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(syscall.Getpid())
	sigchan := make(chan (os.Signal))
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGSEGV, syscall.SIGABRT, syscall.Signal(0x00))
	ticker := time.NewTicker(time.Second)
	tickerBad := time.NewTicker(2 * time.Second)
	doneChan := make(chan (struct{}))
	i := []int{0}
	offset := 0

	defer func() {
		if err := recover(); err != nil {
			// IT DOENS'T WORK
			fmt.Printf("Error came: %v, exiting\n", err)
			ticker.Stop()
			doneChan <- struct{}{}
		}
	}()
loop:
	for {
		select {
		case <-ticker.C:
			i[offset]++
			fmt.Println("Tick number: #%v", i)
		case <-tickerBad.C:

			v := (*int)(unsafe.Pointer(&i))
			x := unsafe.Slice(v, 2)
			x[1] += 1024
			offset += 1024
		case <-doneChan:
			break loop
		case sig := <-sigchan:
			// doesn't work too
			fmt.Printf("Got signal: %v\n", sig)
			break loop

		}
	}
}
