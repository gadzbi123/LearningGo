package main

import (
	"fmt"
	"strings"
)

func pureFuncAppend(fun func(string, string) string) func(string) string {
	s1 := "abc"
	s2 := "dfg"
	return func(s string) string {
		return fun(s1, s2) + s
	}
}
func main() {
	specific := func(s1, s2 string) string {
		return strings.ToUpper(s1) + strings.ToUpper(s2)
	}
	specificAppend := pureFuncAppend(specific)
	s := specificAppend("LOLEX")
	fmt.Println(s)

}
