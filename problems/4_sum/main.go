package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for i := a + 1; i < len(nums)-2; i++ {
			if i > a+1 && nums[i] == nums[i-1] {
				continue
			}
			k, j := i+1, len(nums)-1
			for k < j {
				sum := nums[a] + nums[i] + nums[k] + nums[j]
				if sum > target {
					j--
				} else if sum < target {
					k++
				} else {
					res = append(res, []int{nums[a], nums[i], nums[k], nums[j]})
					lastLow := nums[k]
					lastHigh := nums[j]
					for k < j && nums[k] == lastLow {
						k++
					}
					for k < j && nums[j] == lastHigh {
						j--
					}
				}
			}
		}
	}
	return res
}
