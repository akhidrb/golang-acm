package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return sum
			}
		}
	}
	return closest
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
