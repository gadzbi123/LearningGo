package main

import (
	"fmt"
)

type byteArr []byte

func (arr *byteArr) Write(in []byte) (int, error) {
	// *arr = slices.Clone(in)
	// fmt.Println(string(*arr))
	*arr = in
	// fmt.Println(string(*arr))
	return 5, nil
}
func main() {
	var arr byteArr
	fmt.Fprint(&arr, "I like bananas")
	fmt.Println(len(arr), arr)
	strArr := string(arr)
	fmt.Println(strArr)
}
