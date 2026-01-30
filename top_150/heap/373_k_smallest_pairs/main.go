package main

import "container/heap"

type HeapObj struct {
	sum  int
	i, j int
}

type MinHeap []HeapObj

func (p *MinHeap) Push(val interface{}) {
	*p = append(*p, val.(HeapObj))
}

func (p *MinHeap) Pop() interface{} {
	temp := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return temp
}

func (p *MinHeap) Len() int {
	return len(*p)
}

func (p *MinHeap) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *MinHeap) Less(i, j int) bool {
	return (*p)[i].sum < (*p)[j].sum
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for i := 0; i < len(nums1); i++ {
		sum := nums1[i] + nums2[0]
		heap.Push(minHeap, HeapObj{
			sum: sum,
			i:   i, j: 0,
		})
	}
	res := make([][]int, 0)

	for minHeap.Len() > 0 && k > 0 {
		currMin := heap.Pop(minHeap)
		i := currMin.(HeapObj).i
		j := currMin.(HeapObj).j
		res = append(res, []int{nums1[i], nums2[j]})

		next := j + 1
		if next < len(nums2) {
			sum := nums1[i] + nums2[next]
			heap.Push(minHeap, HeapObj{
				sum: sum,
				i:   i, j: next,
			})
		}
		k--
	}
	return res
}
