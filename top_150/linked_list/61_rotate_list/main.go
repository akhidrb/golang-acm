package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	length := 0
	p1 := head
	for p1 != nil {
		p1 = p1.Next
		length++
	}
	k = k % length

	for i := 0; i < k; i++ {
		bftail, tail := head, head
		if tail.Next != nil {
			tail = tail.Next
		}
		for tail.Next != nil {
			bftail = bftail.Next
			tail = tail.Next
		}
		bftail.Next = nil
		head = &ListNode{Val: tail.Val, Next: head}
	}
	return head
}
