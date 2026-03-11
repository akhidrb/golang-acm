package main

func canJump(nums []int) bool {
	reachable := 0
	for i, val := range nums {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+val)
	}
	return true
}
