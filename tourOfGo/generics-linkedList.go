package main

import (
	"fmt"
)

func findShortest[T comparable](array []T) int {
	var shortestLength = int(^uint(0) >> 1)
	var currIndex = -1
	if len(array) == 0 {
		return currIndex
	}
	for i, x := range array {
		currLen := len(fmt.Sprintf("%s", x))
		if currLen < shortestLength {
			shortestLength = currLen
			currIndex = i
		}
	}
	return currIndex
}

type DoubleLinkedList[T any] struct {
	prev  *DoubleLinkedList[T]
	next  *DoubleLinkedList[T]
	value T
}

func pushBack[T any](list *DoubleLinkedList[T], value T) *DoubleLinkedList[T] {
	if list == nil {
		list = &DoubleLinkedList[T]{nil, nil, value}
		return list
	}
	head := list
	for {
		if list.next == nil {
			break
		}
		list = list.next
	}
	list.next = &DoubleLinkedList[T]{list, nil, value}
	return head
}

func pushFront[T any](list *DoubleLinkedList[T], value T) *DoubleLinkedList[T] {
	if list == nil {
		list = &DoubleLinkedList[T]{nil, nil, value}
		return list
	}
	head := list
	for {
		if list.prev == nil {
			break
		}
		list = list.prev
	}
	list.prev = &DoubleLinkedList[T]{nil, list, value}
	return head
}

func (list *DoubleLinkedList[T]) DisplayFromStart(detailed ...bool) {
	if list == nil {
		fmt.Println("tried to print empty list")
		return
	}
	for {
		if list.prev == nil {
			break
		}
		list = list.prev
	}
	for {
		if len(detailed) == 0 || detailed[0] == false {
			fmt.Printf("%v -> ", list.value)
		} else {
			fmt.Printf("prev:%v, next:%v, ptr:%p, value:%v\n", list.prev, list.next, list, list.value)
		}
		if list.next == nil {
			break
		}
		list = list.next
	}
}

func main() {
	arrInt := []int{2, 4, 77, 124, 2222}
	fmt.Println(findShortest[int](arrInt))
	arrStr := []string{"foo", "bar", "foobaz", "xd"}
	fmt.Println(findShortest[string](arrStr))

	var list *DoubleLinkedList[int]
	list = pushBack[int](list, 5)
	list = pushFront[int](list, 3)
	list = pushFront[int](list, 1)
	list = pushBack[int](list, 6)
	list.DisplayFromStart()
}
