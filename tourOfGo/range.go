package main

import (
	"fmt"
	"strings"
)

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	result = make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
	}
	return result
}

type Person struct {
	name, surrname string
}

func WordCount(s string) map[string]int {
	result := map[string]int{}
	words := strings.Split(s, " ")
	for _, x := range words {
		result[x] += 1
	}
	return result
}

func main() {
	// var people map[string]Person
	people := make(map[string]Person)
	people["Eryk"] = Person{name: "Eryk", surrname: "Grzybek"}
	fmt.Println(people["Eryk"])

	people2 := map[string]Person{"Julian": Person{name: "Julian", surrname: "Borek"}, "Eryk": {name: "Eryk", surrname: "Grzybek"}}
	fmt.Println(people2)
	delete(people2, "Eryk")
	v, ok := people2["Eryk"]

	fmt.Printf("Value of eryk: %v, present? %v", v, ok)

	fmt.Println(WordCount("aaa bbb aaa ccc"))

}
