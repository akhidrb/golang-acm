package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	step := 1
	stack := make([]int, 0, right-left+1)
	p1 := head
	for step < left {
		p1 = p1.Next
		step++
	}
	p2 := p1
	for step <= right {
		stack = append(stack, p1.Val)
		p1 = p1.Next
		step++
	}
	for len(stack) > 0 {
		p2.Val = stack[len(stack)-1]
		p2 = p2.Next
		stack = stack[:len(stack)-1]
	}
	return head
}
