package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{3, 2, 1, 5, 6, 4}
		valid := findKthLargest(nums, 2)
		assert.Equal(t, 5, valid)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{3, 2, 1, 5, 6, 4}
		valid := findKthLargestSorting(nums, 2)
		assert.Equal(t, 5, valid)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
		valid := findKthLargestSorting(nums, 4)
		assert.Equal(t, 4, valid)
	})

}
