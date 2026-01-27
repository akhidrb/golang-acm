package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	output := make([][]int, 0)
	for i, _ := range intervals {
		length := len(output) - 1
		if len(output) == 0 || output[length][1] < intervals[i][0] {
			output = append(output, intervals[i])
		} else {
			output[length][1] = max(output[length][1], intervals[i][1])
		}
	}
	return output
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
