package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// 1, 2, 3, 4, 5
	// fast: 5, slow: 3
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// head: 1, 2, 3
	// p2: 5, 4
	half := slow.Next
	slow.Next = nil
	var p2 *ListNode
	for half != nil {
		temp := half.Next
		half.Next = p2
		p2 = half
		half = temp
	}

	// t1: 2, 3 --- t2: 4 --- p2: 5, 2, 3 --- p1: 1, 5, 2, 3

	p1 := head
	for p2 != nil {
		temp1 := p1.Next
		temp2 := p2.Next
		p2.Next = nil
		p1.Next = p2
		p2.Next = temp1
		p1 = p2.Next
		p2 = temp2
	}
}

func reorderList2(head *ListNode) {
	p1 := head
	p2 := p1
	p3 := p2
	count := 0
	for p2.Next != nil {
		count++
		p2 = p2.Next
		if p2.Next != nil {
			p3 = p3.Next
		}
	}
	count /= 2
	for i := 0; i < count; i++ {
		p3.Next = nil
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p1.Next.Next
		p3 = p2
		for p2.Next != nil {
			p2 = p2.Next
			if p2.Next != nil {
				p3 = p3.Next
			}
		}
	}
}
