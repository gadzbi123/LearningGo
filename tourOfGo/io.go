package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, gadzbi!")

	reader := io.Reader("ABC")
	reader[0]
	b := make([]byte, 20)
	for {
		n, err := r.Read(b)
		fmt.Printf("%v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	fmt.Println(string(b))
}
