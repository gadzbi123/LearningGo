package main

import (
	"fmt"
	"strings"
)

type Sqrter interface {
	Sqrt() []int
}
type MyArr []int

func (arr *MyArr) Sqrt() []int {
	for i, x := range *arr {
		(*arr)[i] *= x
	}
	return *arr
}

type Animal interface {
	make_sound()
}

type Dog struct {
	name  string
	breed string
}

func (d Dog) make_sound() {
	fmt.Printf("Type: %T, %v that has %v breed barks!\n", d, d.name, d.breed)
}

type Cat struct {
	name  string
	color string
}

func (c *Cat) make_sound() {
	fmt.Printf("Type: %T, %v that has %v color meows!\n", c, c.name, strings.ToLower(c.color))
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string
	for i, _ := range ip {
		if i != 0 {
			s += fmt.Sprintf(".")
		}
		s += fmt.Sprintf("%v", ip[i])
	}
	return s
}
func main() {
	var a Sqrter
	myarr := MyArr{1, 2, 3, 4, 5}
	a = &myarr
	myarr.Sqrt()
	fmt.Println(myarr)
	a.Sqrt()
	fmt.Println(a)

	var myAnimal Animal = Dog{name: "Fafik", breed: "Spaniel"}
	myAnimal.make_sound()
	myAnimal = &Cat{name: "Timon", color: "Brown"}
	myAnimal.make_sound()

	var I interface{} = true

	b, ok := I.(bool)
	fmt.Println(b, ok)
	s, ok := I.(string)
	fmt.Println(s, ok)
	//panic
	// f := I.(float64)
	// fmt.Println(f)

	hosts := map[string]IPAddr{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}
	for _, addr := range hosts {
		fmt.Println(addr)
	}
}
