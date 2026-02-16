package main

import "sort"

func sequentialDigits(low int, high int) []int {
	nums := make([]int, 0)
	for i := 1; i <= 9; i++ {
		num := i
		for j := i + 1; j <= 9; j++ {
			num = (num * 10) + j
			if num >= low && num <= high {
				nums = append(nums, num)
			}
		}
	}
	sort.Ints(nums)
	return nums
}
