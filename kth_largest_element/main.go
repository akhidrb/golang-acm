package main

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int { return len(*h) }
func (h *MinHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *MinHeap) Swap(i, j int) {
	temp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = temp
}
func (h *MinHeap) Push(num interface{}) {
	*h = append(*h, num.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]
	return last
}

func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, val := range nums {
		if minHeap.Len() < k {
			heap.Push(minHeap, val)
		} else if val > (*minHeap)[0] {
			(*minHeap)[0] = val
			heap.Fix(minHeap, 0)
		}
	}
	return (*minHeap)[0]
}
