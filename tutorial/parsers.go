package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	if _, err := fmt.Scan(&input); err == nil {

		res, err := strconv.ParseInt(input, 10, 16)
		if numErr, ok := err.(*strconv.NumError); ok {
			if numErr.Err == strconv.ErrRange {
				fmt.Println("range error")
				fmt.Println(input, res)
				return
			}
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(input, res)
		return
	}
}
