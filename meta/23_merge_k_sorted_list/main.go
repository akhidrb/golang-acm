package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(val any) {
	*h = append(*h, val.(*ListNode))
}

func (h *MinHeap) Pop() any {
	n := len(*h)
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}

func mergeKLists2(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	res := &ListNode{}
	curr := res
	for h.Len() > 0 {
		top := heap.Pop(h).(*ListNode)
		curr.Next = top
		curr = curr.Next
		if top.Next != nil {
			top = top.Next
			heap.Push(h, top)
		}
	}
	return res.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return dummy.Next
}

func merge2(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val <= right.Val {
		left.Next = merge(left.Next, right)
		return left
	} else {
		right.Next = merge(left, right.Next)
		return right
	}

}
