package main

import "fmt"

func clean(valid *bool) {
	var out string
	if *valid {
		out = "Is valid"
	} else {
		out = "not valid"
	}
	fmt.Println(out)
}
func main() {
	isValid := true
	defer clean(&isValid)
	for _, v := range []int{1, 2, 3, 4, 5} {
		if v == 6 {
			isValid = false
		}
	}
}
