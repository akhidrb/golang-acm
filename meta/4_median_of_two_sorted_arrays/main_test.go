package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		res := findMedianSortedArrays(nums1, nums2)
		assert.Equal(t, 2.0, res)
	})

	t.Run("2", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		res := findMedianSortedArrays(nums1, nums2)
		assert.Equal(t, 2.5, res)
	})
}
