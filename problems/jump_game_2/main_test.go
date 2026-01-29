package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JumpGame2(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		res := jump(nums)
		assert.Equal(t, 2, res)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int{2, 3, 0, 1, 4}
		res := jump(nums)
		assert.Equal(t, 2, res)
	})
	t.Run("3", func(t *testing.T) {
		nums := []int{2, 1, 1, 1, 4}
		res := jump(nums)
		assert.Equal(t, 3, res)
	})
	t.Run("4", func(t *testing.T) {
		nums := []int{1, 2, 1, 3, 4, 0, 1}
		res := jump(nums)
		assert.Equal(t, 3, res)
	})
}
