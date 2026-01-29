package main

func minimumOperations(nums []int) int {
	h := make(map[int]struct{})
	for _, num := range nums {
		if num > 0 {
			h[num] = struct{}{}
		}
	}
	return len(h)
}
