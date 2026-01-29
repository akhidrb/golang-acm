package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums1 := []int{1, 5, 3, 6, 2, 4}
		res := mergeSort(nums1)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, res)
	})

	t.Run("2", func(t *testing.T) {
		nums1 := []int{1, 5, 7, 3, 6, 2, 4}
		res := mergeSort(nums1)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, res)
	})

}
