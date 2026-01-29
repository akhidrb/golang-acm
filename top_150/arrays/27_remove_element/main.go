package main

func removeElement(nums []int, val int) int {
	k := 0
	for i, num := range nums {
		if num != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
