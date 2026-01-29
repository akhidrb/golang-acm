package main

import (
	"fmt"
	"sort"
	"strings"
)

func threeSum(nums []int) [][]int {
	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
		return [][]int{nums}
	}
	sum := make([][]int, 0)
	uniqueMap := make(map[string]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				data := []int{nums[left], nums[right], nums[i]}
				addToSum(&sum, data, uniqueMap)
				right--
				left++
			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else if nums[left]+nums[right]+nums[i] < 0 {
				left++
			}
		}
	}
	return sum
}

func addToSum(sum *[][]int, data []int, uniqueMap map[string]bool) {
	key := strings.Join(strings.Fields(fmt.Sprint(data)), ",")
	if !uniqueMap[key] {
		uniqueMap[key] = true
		*sum = append(*sum, data)
	}
}
