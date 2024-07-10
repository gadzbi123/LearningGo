package main

import (
	"fmt"
	"strconv"
)

func change(arr []int) {
	fmt.Printf("before in change: %v, %p\n", arr, arr)
	(arr)[1] = 0
	fmt.Printf("after in change: %v, %p\n", arr, arr)
}

func add(arr *[]int) {
	// fmt.Printf("befor in add: %p\n", *arr)
	(*arr) = append(*arr, 5)
	fmt.Printf("after in add: %v\n", cap(*arr))
	// fmt.Printf("after in add: %p\n", *arr)
}

func getStack(x *int) {
	temp := 5
	if x == nil {
		x = &temp
		getStack(x)
	} else {
		y := 5
		fmt.Println(y)
		older := fmt.Sprintf("%#p", x)
		newer := fmt.Sprintf("%#p", &temp)
		fmt.Println(older, newer)
		// older = strings.Replace(older, "0x", "", 1)
		// newer = strings.Replace(newer, "0x", "", 1)
		o, err := strconv.ParseUint(older, 16, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := strconv.ParseUint(newer, 16, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(o, n)
		if o < n {
			fmt.Println("Growing")
		} else {
			fmt.Println("Shrinking")
		}
		return
	}
}
func main() {
	var arr *[]int
	fmt.Printf("before init: %v, %p\n", arr, arr)
	arr = &[]int{5, 3}
	fmt.Printf("after init: %v, %p\n", arr, *arr)
	change(*arr)
	fmt.Printf("after change: %v, %p\n", arr, *arr)
	for i := 0; i < 100; i++ {
		add(arr)
	}
	getStack(nil)
	// fmt.Printf("after add: %v, %p\n", arr, *arr)
}
