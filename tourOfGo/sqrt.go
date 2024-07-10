package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	// var result float64
	var z = 1.0
	for i := 0; ; i++ {
		temp := z
		z -= (z*z - x) / (2 * x)
		if z-temp < .001 {
			fmt.Println("Stopped after %v iterations", i)
			return z
		}
	}
}

func main() {
	result := sqrt(24)
	fmt.Printf("Result is %v", result)
}
