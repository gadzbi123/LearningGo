package main

import (
	"fmt"
	"iter"
	"strings"
)

func filterX(values []string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, v := range values {
			if !strings.Contains(v, "x") && !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	unfilteredVals := []string{"abc", "dfc", "xgg", "dfx", "gg"}
	for _, v := range filterX(unfilteredVals) {
		fmt.Println(v)
	}
}
