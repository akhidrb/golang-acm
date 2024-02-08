package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k, j := i+1, len(nums)-1
		for k < j {
			sum := nums[i] + nums[k] + nums[j]
			if sum > 0 {
				j--
			} else if sum < 0 {
				k++
			} else {
				res = append(res, []int{nums[i], nums[k], nums[j]})
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
	return res
}
