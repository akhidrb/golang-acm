package main

import (
	"container/heap"
)

type MaxHeap []FreqElem

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].freq > m[j].freq
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(FreqElem))
}

func (m *MaxHeap) Pop() any {
	n := len(*m)
	top := (*m)[n-1]
	*m = (*m)[:n-1]
	return top
}

type FreqElem struct {
	num  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	mh := &MaxHeap{}
	heap.Init(mh)
	freq := make(map[int]int)
	for _, value := range nums {
		freq[value]++
	}

	for key, value := range freq {
		freqElem := FreqElem{
			num:  key,
			freq: value,
		}
		heap.Push(mh, freqElem)
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		elem := heap.Pop(mh).(FreqElem)
		res[i] = elem.num
	}
	return res
}
