package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	length := len(citations)
	largest := 0
	for i, val := range citations {
		diff := length - i
		if diff < largest {
			return largest
		}
		if diff <= val {
			largest = diff
		}
	}
	return largest
}
