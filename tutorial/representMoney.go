package main

import (
	"fmt"
	"strconv"
)

func printCash(cash []int) {
	for _, v := range cash {
		str := strconv.Itoa(v)
		switch {

		case len(str) < 3:
			str = fmt.Sprintf("0.%02d", str)
		default:
			str = str[:len(str)-2] + "." + str[len(str)-2:]
		}
		fmt.Println(str)
	}
}
func main() {
	cash := []int{122, 465, 23, 5688, 1}
	printCash(cash)
}
