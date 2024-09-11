package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("errors.go")
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}
	reader := bufio.NewReader(file)
	i := 1
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("exit")
			return
		}
		fmt.Println(i, "->", string(line[:]))
		i++
	}
}
