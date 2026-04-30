package main

func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 && len(nums) >= 2 {
		return true
	}
	count := 0
	two := 0
	for i, val := range nums {
		if i > 0 {
			if two == 2 {
				return true
			}
			if count == k-1 {
				count = 0
				two++
			} else if val > nums[i-1] {
				count++
			} else {
				two = 0
				count = 0
			}
		}
	}
	return false
}
