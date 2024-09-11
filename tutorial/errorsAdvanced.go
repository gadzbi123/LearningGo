package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("first err")
	err2 := fmt.Errorf("error happend: %w", err1)
	fmt.Println(err1)
	fmt.Println(err2)
}
