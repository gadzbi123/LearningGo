package main

import (
	"fmt"
	"math"
)

func main() {
	var my_f float64 = math.Pow(2, 1023)
	var my_i int = int(math.Pow(2, 66))
	fmt.Println(my_f, my_i)
	my_i = int(my_f)
}
