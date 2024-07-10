package test

import (
	"fmt"
	"math"
)

func devide(a int64, b int32) int64 {
	return a / int64(b)
}

func swap(a, b int) (int, int) {
	return b, a
}
func fib(n int) int {
	if n < 2 {
		return 1
	}
	return n + fib(n-1)
}

func pow_with_lim(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}
func testing() {
	fmt.Println(devide(42, 1))
	// var p *int
	// p[1] = 5
	fmt.Println(swap(5, 3))
	var a, b, c = 1, true, "abc"
	k := 5
	fmt.Println(a, b, c, k)
	byting := 12
	var arr []int
	for byting > 0 {
		arr = append([]int{byting & 1}, arr...)
		byting = byting >> 1
	}
	fmt.Printf("%T, %v \n", arr, arr)
	fmt.Printf("%q\n", "")
	var x, y int = 3, 4
	var flut = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(flut)
	fmt.Println(x, y, flut, z)

	var arr1 []int
	for i := 0; i < 23; i++ {
		if i%2 == 0 {
			arr1 = append(arr1, i)
		}
	}
	fmt.Println(arr1)
	fmt.Println(fib(22255))
	fmt.Printf("2 to power of 3 with limit 7 is %v\n", pow_with_lim(2, 3, 7))
	fmt.Printf("2 to power of 3 with limit 16 is %v\n", pow_with_lim(2, 3, 16))
}
