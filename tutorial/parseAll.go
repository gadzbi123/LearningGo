package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Println("Could get a scan", err)
		return
	}
	res, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("res= 0b%b, 0o%o, 0x%x, %v\n", res, res, res, res)
}
