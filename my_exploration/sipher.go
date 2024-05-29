package main

import (
	"fmt"
	"math/rand"
)

func getKey() string {
	key := make([]rune, 0, 256)
	for i := 0; i < 256; i++ {
		random := rand.Int31n(int32(^uint32(0) >> 1))
		key = append(key, random)
	}
	return string(key)
}

func makeSipher(input, key string) string {
	output := make([]rune, 0, len(input))
	in := []rune(input)
	k := []rune(key)
	for i := 0; i < len(in); i++ {
		key_idx := i % 256
		output = append(output, rune(int(k[key_idx])^int(in[i])))
	}
	return string(output)
}

func main() {
	key := getKey()
	var my_string = "lubie w dupę"
	siphered := makeSipher(my_string, key)
	fmt.Println((siphered))
	disiphered := makeSipher(siphered, key)
	fmt.Println((disiphered))
}
