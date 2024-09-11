package main

import (
	"fmt"
	"time"
)

func long_func(i int, start, end chan interface{}) {
	var x interface{}
	start <- x
	fmt.Printf("Running %v task\n", i)
	time.Sleep(time.Second * 2)
	<-start
	end <- x
}
func main() {
	ender := make(chan interface{})
	starter := make(chan interface{}, 2)
	go long_func(1, starter, ender)
	go long_func(2, starter, ender)
	go long_func(3, starter, ender)
	<-ender
	<-ender
	<-ender

}
