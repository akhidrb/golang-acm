package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	length := 0
	for curr != nil {
		length++
		curr = curr.Next
	}
	if n > length {
		return head
	}
	index := length - n
	if length == 1 {
		if index == 0 {
			head = nil
		}
		return head
	}
	curr = head.Next
	prev := head
	i := 1
	if i > index {
		return head.Next
	}
	for curr != nil {
		if i == index {
			prev.Next = curr.Next
			break
		}
		i++
		curr = curr.Next
		prev = prev.Next
	}
	return head
}
