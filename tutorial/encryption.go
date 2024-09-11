package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	message := "Lubie w dupe"
	res := sha256.Sum256([]byte(message))
	messageAlterred := "Lubie w pupe"
	resAltered := sha256.Sum256([]byte(messageAlterred))
	fmt.Printf("sha 1:%x\n", (res))
	fmt.Printf("sha 2:%x\n", (resAltered))
	res1 := md5.Sum([]byte(message))
	res2 := md5.Sum([]byte(messageAlterred))
	fmt.Printf("md5 1:%x\n", res1)
	fmt.Printf("md5 2:%x\n", res2)
}
