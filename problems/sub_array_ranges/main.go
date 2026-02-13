package main

func subArrayRanges(nums []int) int64 {
	sum := 0
	for i := 0; i < len(nums); i++ {
		largest := nums[i]
		smallest := nums[i]
		for j := i + 1; j < len(nums); j++ {
			largest = max(largest, nums[j])
			smallest = min(smallest, nums[j])
			sum += largest - smallest
		}
	}
	return int64(sum)
}
