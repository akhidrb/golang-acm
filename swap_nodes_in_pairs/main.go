package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	isFirst := true
	p1 := head
	p2 := head
	p3 := head.Next
	for {
		p2.Next = p3.Next
		p3.Next = p2
		if isFirst {
			head = p3
			isFirst = false
		} else {
			p1.Next = p3
		}
		if p2.Next != nil && p2.Next.Next != nil {
			p1 = p2
			p3 = p2.Next.Next
			p2 = p2.Next
			continue
		}
		break
	}
	return head
}
