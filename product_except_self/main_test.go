package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ProductExceptSelf(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		res := productExceptSelf(nums)
		assert.Equal(t, []int{24, 12, 8, 6}, res)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		res := productExceptSelf(nums)
		assert.Equal(t, []int{0, 0, 9, 0, 0}, res)
	})
}
