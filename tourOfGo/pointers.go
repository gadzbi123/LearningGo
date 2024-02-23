package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func main() {
	i := 42
	p := &i
	fmt.Printf(" %T %v\n", p, *p)
	my_point := Point{X: 2.}
	my_point.Y = 3.
	point_pointer := &my_point
	point_pointer.X = 3e-12
	fmt.Printf("%v\n", *point_pointer)
	fmt.Println(my_point)

	SIZE := 3
	primes := [6]int{2, 3, 5, 7, 11, 13}
	inner3 := make([]int, SIZE)
	prime_copy := primes
	copy(inner3, primes[1:4])
	inner3[2] = 55
	prime_copy[2] = 12
	fmt.Println(inner3)
	fmt.Println(primes)
	fmt.Println(prime_copy)
	part := prime_copy[3:]
	part2 := prime_copy[:2]
	part12 := append(part, part2...)
	inner3 = append(inner3, part12...)
	fmt.Println(inner3)
}
