package main

import "math/rand/v2"

type Solution struct {
	prefixSum []int
	sum       int
}

func Constructor(w []int) Solution {
	sum := 0
	prefixSum := make([]int, len(w))
	for i, val := range w {
		sum += val
		prefixSum[i] = sum
	}
	return Solution{prefixSum: prefixSum, sum: sum}
}

func (this *Solution) findIndex(n int) int {
	start, end := 0, len(this.prefixSum)
	for start < end {
		mid := start + (end-start)/2
		if this.prefixSum[mid] < n {
			start = mid + 1
		} else if this.prefixSum[mid] > n {
			end = mid - 1
		} else {
			return mid
		}
	}
	return start
}

func (this *Solution) PickIndex() int {
	ranNum := rand.IntN(this.sum) + 1
	return this.findIndex(ranNum)
}
