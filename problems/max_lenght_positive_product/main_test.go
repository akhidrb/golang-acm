package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, -2, -3, 4}
		res := getMaxLen(nums)
		assert.Equal(t, 4, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 1, -2, -3, -4}
		res := getMaxLen(nums)
		assert.Equal(t, 3, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{-1, -2, -3, 0, 1}
		res := getMaxLen(nums)
		assert.Equal(t, 2, res)
	})

}
