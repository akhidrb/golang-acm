package main

func quicksort(nums []int) {
	qs(nums, 0, len(nums)-1)
}

func qs(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(nums, left, right)
	qs(nums, left, pivotIndex-1)
	qs(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	i := left
	pivot := nums[right]

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[right], nums[i] = nums[i], nums[right]
	return i
}
