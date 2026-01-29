package main

import "sort"

func removeElement(nums []int, val int) int {
	k := 0
	for i, num := range nums {
		if num == val {
			nums[i] = -1
		} else {
			k++
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return k
}
