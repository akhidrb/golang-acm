package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	left := sortList(head)
	right := sortList(mid)

	return merge(left, right)
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for left != nil && right != nil {
		if right.Val < left.Val {
			tail.Next = right
			right = right.Next
		} else {
			tail.Next = left
			left = left.Next
		}
		tail = tail.Next
	}
	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}
	return head.Next
}
