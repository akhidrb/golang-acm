package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{-2, 0, 2, -2, 1, -1}
		sums := threeSum(nums)
		assert.Equal(t, [][]int{{-2, 0, 2}, {-1, 0, 1}}, sums)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		sums := threeSum(nums)
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, sums)
	})
}
