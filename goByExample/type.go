package main

import "fmt"

func main() {
	var myType interface{}
	myType = uint64(6)
	switch x := myType.(type) {
	default:
		fmt.Printf("Type: %T", uint16(x.(uint64))+5)
	}
}
