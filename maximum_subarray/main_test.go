package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		res := maxSubArray(nums)
		assert.Equal(t, 6, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1}
		res := maxSubArray(nums)
		assert.Equal(t, 1, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{5, 4, -1, 7, 8}
		res := maxSubArray(nums)
		assert.Equal(t, 23, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{-2, 1}
		res := maxSubArray(nums)
		assert.Equal(t, 1, res)
	})

}
