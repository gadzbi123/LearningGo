package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func IsCyclic(head *ListNode) bool {
	if head == nil {
		return false
	}
	temp1 := head.Next
	temp2 := head
	for temp1 != nil {
		for {
			if temp1 == temp2 || temp1.Next == temp2 {
				return true
			}
			temp2 = temp2.Next
			if temp2 == temp1 {
				temp2 = head
				break
			}
		}
		temp1 = temp1.Next
	}
	return false
}
func arrToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head, temp, prev_temp *ListNode
	for i := 0; i < len(arr); i++ {
		temp = &ListNode{Value: arr[i], Next: nil}
		if prev_temp != nil {
			prev_temp.Next = temp
		}
		if head == nil {
			head = temp
		}
		prev_temp = temp
		temp = temp.Next
	}
	return head
}

func arrToListCyclic(arr []int, tailPointsTo int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head, temp, prev_temp, tail *ListNode
	for i := 0; i < len(arr); i++ {
		temp = &ListNode{Value: arr[i], Next: nil}
		if prev_temp != nil {
			prev_temp.Next = temp
		}
		if head == nil {
			head = temp
		}
		if tailPointsTo == i {
			tail = temp
		}
		if i+1 == len(arr) {
			temp.Next = tail
		}
		prev_temp = temp
		temp = temp.Next
	}
	return head
}

func (head *ListNode) printLimit(size int) {
	var counter int
	for head != nil && counter < size {
		fmt.Printf("(%p,{%v,%p}),", head, head.Value, head.Next)
		head = head.Next
		counter++
	}
	fmt.Println()
}

func (head *ListNode) print() {
	for head != nil {
		fmt.Printf("%v,", head.Value)
		head = head.Next
	}
	fmt.Println()
}
func taskIsCyclic() {
	list := arrToListCyclic([]int{1, 2, 3, 4, 5}, -1)
	list.printLimit(5)
	fmt.Println(IsCyclic(list))

	list = arrToListCyclic([]int{3, 2, 0, -4}, 1)
	list.printLimit(5)
	fmt.Println(IsCyclic(list))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, res, temp *ListNode
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		var sum int
		if l1 != nil {
			sum += l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Value
			l2 = l2.Next
		}
		sum += carry
		rem := sum % 10
		carry = sum / 10
		res = &ListNode{Value: rem, Next: nil}
		if head == nil {
			head = res
		}
		if temp != nil {
			temp.Next = res
		}
		temp = res
		res = res.Next
	}
	return head
}
func main() {
	list1 := arrToList([]int{2, 4, 3})
	list2 := arrToList([]int{5, 6, 4})
	res := addTwoNumbers(list1, list2)
	res.print()
	res = addTwoNumbers(arrToList([]int{9, 9, 9, 9, 9, 9, 9}), arrToList([]int{9, 9, 9, 9}))
	res.print()
}
