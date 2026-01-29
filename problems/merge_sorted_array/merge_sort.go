package main

func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	first := mergeSort(arr[:n/2])
	last := mergeSort(arr[n/2:])
	return merge2(first, last)
}

func merge2(nums1, nums2 []int) []int {
	nums3 := make([]int, len(nums1)+len(nums2))
	for i, val := range nums1 {
		nums3[i] = val
	}
	merge(nums3, len(nums1), nums2, len(nums2))
	return nums3
}
