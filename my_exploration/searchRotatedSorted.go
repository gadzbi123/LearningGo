package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func search(nums []int, target int) int {
	mid := len(nums) / 2
	pivot := -1
	for {
		m := mid / 2
		if Abs(m-mid) == 1 {
			pivot = min(m+mid, mid)
			return pivot
		}
		if nums[mid] < nums[m] {
			fmt.Println("left")
			mid = m
		} else {
			fmt.Println("right")
			mid = m + mid
		}
	}

}
func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))

}
