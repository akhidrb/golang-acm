package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	var rem int
	head := &ListNode{}
	l3 := head
	for p1 != nil || p2 != nil || rem > 0 {
		val1, val2 := 0, 0
		if p1 != nil {
			val1 = p1.Val
		}
		if p2 != nil {
			val2 = p2.Val
		}
		sum := val1 + val2 + rem
		rem = sum / 10
		sum = sum % 10
		l3.Val = sum
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
		if p1 != nil || p2 != nil || rem > 0 {
			l3.Next = &ListNode{}
			l3 = l3.Next
		}
	}
	return head
}

func main() {
	list1 := []int{2, 4, 9}
	list2 := []int{5, 6, 4, 9}
	l1 := &ListNode{}
	p1 := l1
	for i, val := range list1 {
		p1.Val = val
		if i < len(list1)-1 {
			p1.Next = &ListNode{}
		}
		p1 = p1.Next
	}
	l2 := &ListNode{}
	p2 := l2
	for i, val := range list2 {
		p2.Val = val
		if i < len(list2)-1 {
			p2.Next = &ListNode{}
		}
		p2 = p2.Next
	}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
