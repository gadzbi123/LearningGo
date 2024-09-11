package main

import (
	"errors"
	"fmt"
	"unsafe"
)

const nameSize = 10
const pSize = unsafe.Sizeof(Person{})
const memPoolSize = 1000

var memPool [memPoolSize]byte
var SP uintptr = 0

type Person struct {
	age  uint8
	name [nameSize]byte
}

func allocOnMemPool(p Person) {
	pPointer := *(*[pSize]byte)(unsafe.Pointer(&p))
	copy(memPool[SP:SP+pSize+1], pPointer[:])
	SP += pSize
}

func getFromMemPool(ptr uintptr) (p Person, err error) {
	if ptr+pSize > memPoolSize {
		return p, errors.New("Person is outside bounds")
	}

	age := *(*uint8)(unsafe.Pointer(&memPool[ptr]))
	nameBytes := (unsafe.Slice(&memPool[ptr+1], nameSize))
	var name [nameSize]byte
	copy(name[:], nameBytes)
	return Person{age: age, name: name}, nil
}

func initName(name string) (nameArr [nameSize]byte, err error) {
	nameBytes := []byte(name)
	if len(nameBytes) > nameSize {
		return [nameSize]byte{}, errors.New(fmt.Sprintf("Name %v is longer then max: %v", name, nameSize))

	}
	for i, c := range nameBytes[:] {
		nameArr[i] = c
	}
	return
}

func main() {
	nameArr, err := initName("Tomas")
	p := Person{age: 10, name: [nameSize]byte(nameArr)}
	allocOnMemPool(p)
	fmt.Printf("%v\n", memPool[:20])
	nameArr, err = initName("Bogdan")
	if err != nil {
		panic("Name to long")
	}
	p = Person{age: 20, name: [nameSize]byte(nameArr)}
	allocOnMemPool(p)
	fmt.Printf("%v\n", memPool[:20])
	p, err = getFromMemPool(11)
	fmt.Println(p.age, string(p.name[:]))
}
