package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums1 := []int{4, 1, 2}
		nums2 := []int{1, 3, 4, 2}
		res := nextGreaterElement(nums1, nums2)
		exp := []int{-1, 3, -1}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums1 := []int{2, 4}
		nums2 := []int{1, 2, 3, 4}
		res := nextGreaterElement(nums1, nums2)
		exp := []int{3, -1}
		assert.Equal(t, exp, res)
	})

}
