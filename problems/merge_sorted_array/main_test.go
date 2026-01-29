package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSortedArray(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		m, n := 3, 3
		merge(nums1, m, nums2, n)
		assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
	})

	t.Run("2", func(t *testing.T) {
		nums1 := []int{1}
		nums2 := make([]int, 0)
		m, n := 1, 0
		merge(nums1, m, nums2, n)
		assert.Equal(t, []int{1}, nums1)
	})

	t.Run("3", func(t *testing.T) {
		nums1 := []int{0}
		nums2 := []int{1}
		m, n := 0, 1
		merge(nums1, m, nums2, n)
		assert.Equal(t, []int{1}, nums1)
	})

}
