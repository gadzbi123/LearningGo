package main

import "fmt"

func main() {
	my_map := map[string]any{
		"value": 2.5,
	}
	var x = 5
	v, ok := my_map["value"].(int)
	if !ok {
		fmt.Println("value was not int")
		return
	}
	y := x + v
	fmt.Println("Result", y)

}
