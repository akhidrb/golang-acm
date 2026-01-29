package main

import "sort"

func longestConsecutive(nums []int) int {
	largest := 0
	iter := 1
	if len(nums) == 0 {
		return largest
	}
	consMap := make(map[int]interface{})
	keys := make([]int, 0)
	for _, val := range nums {
		if _, ok := consMap[val]; !ok {
			keys = append(keys, val)
		}
		consMap[val] = nil
	}
	sort.Ints(keys)
	for _, val := range keys {
		if _, ok := consMap[val-1]; ok {
			iter++
		} else {
			if iter > largest {
				largest = iter
			}
			iter = 1
		}
	}
	if iter > largest {
		largest = iter
	}
	return largest
}
