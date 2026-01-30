package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums1 := []int{1, 7, 11}
		nums2 := []int{2, 4, 6}
		k := 3
		res := kSmallestPairs(nums1, nums2, k)
		resExp := [][]int{{1, 2}, {1, 4}, {1, 6}}
		assert.Equal(t, resExp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums1 := []int{1, 1, 2}
		nums2 := []int{1, 2, 3}
		k := 2
		res := kSmallestPairs(nums1, nums2, k)
		resExp := [][]int{{1, 1}, {1, 1}}
		assert.Equal(t, resExp, res)
	})

}
