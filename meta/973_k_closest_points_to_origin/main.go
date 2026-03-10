package main

import (
	"container/heap"
)

type Points struct {
	x   int
	y   int
	sum int
}

type MaxHeap []Points

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].sum > h[j].sum
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Points))
}

func (h *MaxHeap) Pop() any {
	n := len(*h)
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < len(points); i++ {
		p1, p2 := points[i][0], points[i][1]
		sum := (p1 * p1) + (p2 * p2)
		heap.Push(h, Points{
			x:   p1,
			y:   p2,
			sum: sum,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([][]int, k)
	i := 0
	for h.Len() > 0 {
		top := heap.Pop(h).(Points)
		res[i] = []int{top.x, top.y}
		i++
	}
	return res
}
