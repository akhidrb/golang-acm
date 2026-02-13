package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	indexMap := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		indexMap[nums2[i]] = i
		top := -1
		if len(stack) > 0 {
			top = stack[len(stack)-1]
			for top <= nums2[i] {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				} else {
					top = -1
					break
				}
			}
		}
		stack = append(stack, nums2[i])
		nums2[i] = top
	}
	for i := 0; i < len(nums1); i++ {
		index := indexMap[nums1[i]]
		nums1[i] = nums2[index]
	}
	return nums1
}
