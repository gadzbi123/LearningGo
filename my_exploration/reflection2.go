package main

import (
	"fmt"
	"reflect"
	"slices"
)

func (t T) FindAndRemove(slice T, element U) (result T) {
	// el := element.Elem()
	length := len(slice)
	x := -1
	for i := 0; i < length; i++ {
		if slice[i] == element {
			x = i
		}
	}
	if x == -1 {
		result = slice
	} else {
		result = slices.Concat(slice[:x], slice[x+1:])
	}
	return result
}

type U int
type T []U

func main() {
	ints := T{1, 2, 3, 4, 5}
	fn := reflect.ValueOf(ints).MethodByName("FindAndRemove")
	rf := []reflect.Value{reflect.ValueOf(ints), reflect.ValueOf(U(3))}
	res := fn.Call(rf)
	fmt.Println(res[0])
}
