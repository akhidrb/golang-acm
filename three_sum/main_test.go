package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		res := threeSum(nums)
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 1, 1}
		res := threeSum(nums)
		assert.Equal(t, [][]int{}, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{0, 0, 0}
		res := threeSum(nums)
		assert.Equal(t, [][]int{{0, 0, 0}}, res)
	})

}
