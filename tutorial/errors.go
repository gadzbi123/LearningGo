package main

import "fmt"

func division(a, b float64, res chan float64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if b == 0 {
		panic("Divider should not be zero")
	}
	res <- a / b
}
func main() {
	results := make(chan float64)
	scenario := [][2]float64{[2]float64{2, 4}, [2]float64{2, 6}, [2]float64{2, 0}, [2]float64{2, 3}}
	for _, s := range scenario {
		go division(s[0], s[1], results)
	}
	for i := 0; i < len(scenario); i++ {
		fmt.Println("Result:", <-results)
	}
}
