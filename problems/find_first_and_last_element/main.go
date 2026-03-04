package main

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	found := false
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			found = true
			break
		}
	}
	if !found {
		return []int{-1, -1}
	}
	l, r = mid, mid
	for r < len(nums)-1 && nums[r+1] == target {
		r++
	}
	for l > 0 && nums[l-1] == target {
		l--
	}
	return []int{l, r}
}
