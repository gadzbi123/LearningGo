package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Bąk"
	lngByte := len(str)
	lngRune := utf8.RuneCountInString(str)
	fmt.Println("Len of str:", lngRune)
	fmt.Println("Len of bytes:", lngByte)
}
