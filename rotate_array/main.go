package main

func rotate(nums []int, k int) {
	if len(nums) == 1 || len(nums)-k == 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	temp := make([]int, len(nums))
	l, r := 0, len(nums)-k
	for l < len(nums)-k || r < len(nums) {
		if r < len(nums) {
			temp[l] = nums[r]
			r++
		}
		if l < len(nums)-k {
			temp[l+k] = nums[l]
		}
		l++
	}
	copy(nums, temp)
}
