package main

func nextPermutation(nums []int) {
	n := len(nums)
	idx := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx = i
			break
		}
	}
	if idx == -1 {
		reverse(nums, 0, n-1)
		return
	}
	reverse(nums, idx+1, n-1)
	newJ := -1
	for j := idx + 1; j < n; j++ {
		if nums[idx] < nums[j] {
			newJ = j
			break
		}
	}
	nums[newJ], nums[idx] = nums[idx], nums[newJ]
}

func reverse(nums []int, start, end int) {
	p1, p2 := start, end
	for p1 < p2 {
		nums[p1], nums[p2] = nums[p2], nums[p1]
		p1++
		p2--
	}
}
