package main

func removeDuplicates(nums []int) int {
	pivotVal := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != pivotVal {
			pivotVal = nums[i]
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
