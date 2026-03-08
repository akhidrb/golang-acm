package main

import "fmt"

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	qs(nums, 0, len(nums)-1)
}

func qs(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition(nums, left, right)
	fmt.Println(nums)
	qs(nums, left, pivotIndex-1)
	qs(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right] // choose last element as pivot
	i := left            // boundary of smaller elements

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// place pivot in its correct position
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {
	nums := []int{5, 3, 8, 4, 2, 7, 1, 10}
	quickSort(nums)
	fmt.Println(nums)
}
