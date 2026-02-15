package main

func getMaxLen(nums []int) int {
	maxLen := 0
	pos := make([]int, len(nums)+1)
	neg := make([]int, len(nums)+1)
	neg[0] = 0
	pos[0] = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			neg[i+1] = 1 + pos[i]
			if neg[i] > 0 {
				pos[i+1] = 1 + neg[i]
			}
		} else if nums[i] == 0 {
			pos[i+1] = 0
			neg[i+1] = 0
		} else {
			pos[i+1] = 1 + pos[i]
			if neg[i] > 0 {
				neg[i+1] = 1 + neg[i]
			}
		}
		maxLen = max(pos[i+1], maxLen)
	}

	return maxLen
}
