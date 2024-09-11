package main

import (
	"fmt"
)

func Printfln(template string, values ...any) (int, error) {
	return fmt.Printf(template+"\n", values...)
}

func main() {
	var cat, name string
	var price float64
	fmt.Println("Write category, name and price")
	n, err := fmt.Scanln(&cat, &name, &price)
	if err == nil {
		Printfln("Size: %d", n)
		Printfln("%v, %v, %.2f", cat, name, price)
	} else {
		Printfln("Error: %v", err.Error())
	}

}
