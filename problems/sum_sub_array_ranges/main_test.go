package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		res := subArrayRanges(nums)
		assert.Equal(t, int64(4), res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 3, 3}
		res := subArrayRanges(nums)
		assert.Equal(t, int64(4), res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{4, -2, -3, 4, 1}
		res := subArrayRanges(nums)
		assert.Equal(t, int64(59), res)
	})

}
