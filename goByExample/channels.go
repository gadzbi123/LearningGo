package main

import (
	"fmt"
	"strings"
	"time"
)

func toUpper(ch chan string) {
	x := <-ch
	time.Sleep(time.Second)
	ch <- strings.ToUpper(x)
}

func toUpperDouble(out chan<- string, in <-chan string) {
	x := <-in
	time.Sleep(time.Second)
	out <- strings.ToUpper(x)
}
func main() {
	shared := make(chan string)
	sender, reciver := make(chan string), make(chan string)
	str := "Lubie placki"

	go toUpperDouble(reciver, sender)
	go toUpper(shared)
	func() {
		sender <- str
		shared <- str
	}()
	// fmt.Printf("Gave: %q, recived: %q\n", str, <-shared)
	// fmt.Println(<-reciver)
	for i := 0; i < 2; i++ {
		select {
		case x := <-reciver:
			fmt.Println(x)
		case x := <-shared:
			fmt.Printf("Gave: %q, recived: %q\n", str, x)
		}
	}

}
