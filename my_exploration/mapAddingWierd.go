package main

import "fmt"

func main() {
	m := map[int]int{0: 0, 1: 100, 2: 200}
	r, n, i := len(m), len(m), 0
	for range m {
		m[n] = n * 100
		n++
		i++
	}
	fmt.Printf("%d new entries, iterate %d and skip %d\n",
		i, i-r, n-i,
	)
	fmt.Println(m)
}
