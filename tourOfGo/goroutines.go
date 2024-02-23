package main

import (
	"fmt"
	"math/rand"
)

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		c <- y
		x, y = y, x+y
	}
	close(c)
}

func fibControlledByFunc() {
	c := make(chan int, 10)
	go fib(cap(c), c)
	for i := range c {
		fmt.Printf("%v: %v \n", i, <-c)
	}
	fmt.Println()
}

func fibDoubleComs(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibControlledByMaster() {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%v: %v \n", i, <-c)
		}
		quit <- true
	}()
	fibDoubleComs(c, quit)
}

type primeResult struct {
	prime   int
	isPrime bool
}

func isPrime(n int, c chan primeResult) bool {
	if n < 2 {
		c <- primeResult{n, false}
		return false
	}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			c <- primeResult{n, false}
			return false
		}
	}
	c <- primeResult{n, true}
	return true
}

func readPrimesSelect(results_channel chan primeResult, quit_chan chan bool) {
	for {
		select {
		case res := <-results_channel:
			fmt.Println(res)
		case isQuit := <-quit_chan:
			if isQuit {
				return
			}
		default:
		}
	}
}
func readPrimesIf(results_channel chan primeResult, quit_chan chan bool) {
	for {
		res := <-results_channel
		fmt.Println(res)
		if <-quit_chan {
			return
		}
	}
}
func generatePrimes(results_channel chan primeResult, quit_chan chan bool) {
	for {
		v := isPrime(rand.Int()%100, results_channel)
		quit_chan <- v
		if v {
			return
		}
	}
}
func main() {
	results_channel := make(chan primeResult)
	quit_chan := make(chan bool)
	go readPrimesIf(results_channel, quit_chan)
	generatePrimes(results_channel, quit_chan)
}
