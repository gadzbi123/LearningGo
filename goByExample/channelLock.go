package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return 1
	}
	prev := 1
	res := 1
	for i := 1; i < n; i++ {
		temp := res
		res += prev
		prev = temp
	}
	return res
}

type Result struct {
	result, goroutine int
}

func process(res chan Result, fibNum, goroutine int) {
	res <- Result{fib(fibNum), goroutine}
}
func main() {
	result := make(chan Result, 1)
	for i := 0; i < 10; i++ {
		if i == 5 {
			go process(result, 10, i)
		} else {
			go process(result, 1000, i)
		}
	}
	res := <-result
	fmt.Printf("Result: %v from routine %v\n", res.result, res.goroutine)
	fmt.Println("Shit is so unreliable, fuck that")
}
