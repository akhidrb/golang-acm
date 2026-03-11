package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2 := 0, 0
	size := len(nums1) + len(nums2)
	nums3 := make([]int, 0, size)
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] <= nums2[p2] {
			nums3 = append(nums3, nums1[p1])
			p1++
		} else {
			nums3 = append(nums3, nums2[p2])
			p2++
		}
	}
	nums3 = append(nums3, nums1[p1:]...)
	nums3 = append(nums3, nums2[p2:]...)
	if size == 1 {
		return float64(nums3[0])
	}
	if size%2 == 0 {
		ind := size / 2
		return (float64(nums3[ind]) + float64(nums3[ind-1])) / 2.0
	}
	return float64(nums3[size/2])
}
