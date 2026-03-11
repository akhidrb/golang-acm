package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	list3 := &ListNode{}
	p3 := list3
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			list3.Next = p1
			p1 = p1.Next
		} else {
			list3.Next = p2
			p2 = p2.Next
		}
		list3 = list3.Next
	}
	for p1 != nil {
		list3.Next = p1
		p1 = p1.Next
		list3 = list3.Next
	}
	for p2 != nil {
		list3.Next = p2
		p2 = p2.Next
		list3 = list3.Next
	}

	return p3.Next
}
