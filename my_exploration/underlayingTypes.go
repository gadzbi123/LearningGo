package main

import "fmt"

type MyInt int64
type TypeA *int64
type TypeB *MyInt

func makePtr[T any](value T) *T {
	return &value
}
func main() {
	var a TypeA = makePtr(int64(5))
	var b TypeB = (*MyInt)(makePtr(int64(10)))

	x := (*MyInt)(b)
	a = (*int64)(x)
	fmt.Println(*a)
	/*
		x := (*int64)((*MyInt)(b))
		a = x
	*/
}
