package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	arr := make([][]int, 0)
	for _, interval := range intervals {
		if len(arr) == 0 || interval[0] > arr[len(arr)-1][1] {
			arr = append(arr, interval)
		} else {
			arr[len(arr)-1][1] = max(interval[1], arr[len(arr)-1][1])
		}
	}
	return arr
}
