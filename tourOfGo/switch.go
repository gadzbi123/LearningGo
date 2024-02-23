package main

import (
	"fmt"
	"log"
	"time"
)

func devide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't devide by 0")
	}
	return a / b, nil
}

func main() {
	second := time.Now().Second() % 10
	switch second {
	case 1:
		fmt.Println("First second")
	case 2:
		fmt.Println("Second second")
	case 3:
		fmt.Println("Third second")
	default:
		fmt.Printf("%v second\n", second)
	}
	value, err := devide(1, 0)
	if err != nil {
		log.SetFlags(30)
		x := log.Flags()
		fmt.Println(x)
		return
	}
	fmt.Println("Result from devision: ", value)
}
