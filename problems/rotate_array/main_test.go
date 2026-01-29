package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RotateArray(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		rotate(nums, k)
		assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int{-1, -100, 3, 99}
		k := 2
		rotate(nums, k)
		assert.Equal(t, []int{3, 99, -1, -100}, nums)
	})
	t.Run("3", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 2
		rotate(nums, k)
		assert.Equal(t, []int{2, 3, 1}, nums)
	})
}
