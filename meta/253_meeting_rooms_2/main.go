package main

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	n := len(*h)
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &MinHeap{}
	heap.Init(h)
	for _, interval := range intervals {
		if h.Len() > 0 && interval[0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}
	return h.Len()
}
