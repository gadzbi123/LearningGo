package main

import "fmt"

func maxSubArray(nums []int) int {
	best := ^int(^uint(1 << 63))
	curr := 0
	for _, x := range nums {
		curr = max(x, curr+x)
		best = max(curr, best)
	}
	return best

}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
