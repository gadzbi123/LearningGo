package main

import (
	"bytes"
	"fmt"
)

func zigzag(str []byte, n int) []byte {
	i := 0
	up := false
	curr_char := 0
	storage := make([][]byte, n)
	for curr_char != len(str) {
		storage[i] = append(storage[i], str[curr_char])
		curr_char++
		// if i+1 == n {
		// 	up = true
		// }
		// if i == 0 {
		// 	up = false
		// }
		if !up {
			i++
		} else {
			i--
		}
		if i+1 == n || i == 0 {
			up = !up
		}
	}
	return bytes.Join(storage, []byte(""))
}
func main() {
	pattern := []byte("ABCDEFG")
	out := zigzag(pattern, 3)
	fmt.Println(string(out))
}
