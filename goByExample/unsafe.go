package main

import (
	"fmt"
	"unsafe"
)

type Block struct {
	a int
	b bool
	c int
}
type BoolBlock struct {
	a bool
	b bool
	c bool
}
type StrBlock struct {
	a string
	b bool
	c string
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5}
	elemSize := int(unsafe.Sizeof(arr[0]))
	fmt.Println(elemSize)
	last := unsafe.Pointer(&arr[4])
	prevLast := (*int)(unsafe.Add(last, -1*elemSize))
	fmt.Println(*prevLast)
	arrOfx := unsafe.Slice(prevLast, 7)
	fmt.Println(arrOfx)
	arrOfx[6] = 33
	fmt.Println(arrOfx)

	b := Block{0, true, 2}
	pb := unsafe.Pointer(&b)
	fmt.Println(unsafe.Offsetof(b.a))
	fmt.Println(unsafe.Offsetof(b.b))
	fmt.Println(unsafe.Offsetof(b.c))
	wtf := (*int)(unsafe.Add(pb, 8*2))
	fmt.Println("value from struct", wtf, *wtf)
	for i := 1; i < 10; i++ {
		outside := (*int)(unsafe.Add(pb, -8*i))
		fmt.Printf("value outside struct: %v, %v, %v\n", i, outside, *outside)
	}

	bb := BoolBlock{false, true, false}
	pbb := unsafe.Pointer(&bb)
	fmt.Println(unsafe.Offsetof(bb.a))
	fmt.Println(unsafe.Offsetof(bb.b))
	fmt.Println(unsafe.Offsetof(bb.c))
	wtf2 := (*bool)(unsafe.Add(pbb, 2))
	fmt.Println("value from struct", wtf2, *wtf2)
	for i := -1; i > -10; i-- {
		outside := (*bool)(unsafe.Add(pbb, i))
		fmt.Printf("value outside struct: %v, %v, %v\n", i, outside, *outside)
	}

	sb := StrBlock{"abc", true, "ccccc"}
	psb := unsafe.Pointer(&sb)
	fmt.Println(unsafe.Sizeof(sb))
	fmt.Println(unsafe.Offsetof(sb.a))
	fmt.Println(unsafe.Offsetof(sb.b))
	fmt.Println(unsafe.Offsetof(sb.c))
	wtf3 := (*bool)(unsafe.Add(psb, 16))
	fmt.Println("value from struct", wtf3, *wtf3)
	wtf4 := (*string)(unsafe.Add(psb, 24))
	fmt.Println("value from struct", wtf4, *wtf4)
	strPtr := (*string)(unsafe.Add(psb, 0))
	fmt.Println("value from struct", strPtr, *strPtr)
	size := (*int)(unsafe.Add(psb, 8))
	fmt.Println("value from struct", size, *size)
	for i := 0; i < 5; i++ {
		outside := (*uint8)(unsafe.Add(psb, 8*i))
		fmt.Printf("value outside struct: %v, %v, %v\n", i, outside, *outside)
	}
}
