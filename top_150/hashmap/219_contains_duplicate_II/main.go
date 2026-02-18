package main

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		p := nums[i]
		if val, ok := dict[p]; ok {
			if i-val <= k {
				return true
			}
		}
		dict[p] = i
	}
	return false
}
