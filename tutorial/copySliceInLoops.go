package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for range s {
		s = append(s, 10)
		fmt.Printf("inloop s: %v\n", s)
	}
	fmt.Printf("outloop s: %v\n", s)

	s = []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
		fmt.Printf("inloop s: %v\n", s)
		if len(s) > 10 {
			fmt.Printf("TOO MANY VALUES IN s\n")
			break
		}
	}
	fmt.Printf("outloop s: %v\n", s)
}
