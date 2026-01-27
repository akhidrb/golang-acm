package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	numsM := make([]int, m)
	copy(numsM, nums1[:m])
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if numsM[i] < nums2[j] {
			nums1[k] = numsM[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for i < m {
		nums1[k] = numsM[i]
		i++
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}
