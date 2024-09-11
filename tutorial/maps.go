package main

import "fmt"

func addType(x map[string]uint, k string, v uint) {
	x[k] = v
}
func main() {
	var types map[string]uint = map[string]uint{"bool": 1, "byte": 8, "uint": 32, "uint64": 64}
	addType(types, "joke4", 4)
	for k, v := range types {
		fmt.Println(k, v)
	}
	if wierd, ok := types["joke3"]; ok {
		fmt.Println(wierd)
	}

}
