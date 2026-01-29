package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TrappingRainWater(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		res := trap(nums)
		assert.Equal(t, 6, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{4, 2, 0, 3, 2, 5}
		res := trap(nums)
		assert.Equal(t, 9, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{3, 0, 2, 0, 5}
		res := trap(nums)
		assert.Equal(t, 7, res)
	})

}
