package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		res := findPeakElement(nums)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 2, 1, 3, 5, 6, 4}
		res := findPeakElement(nums)
		assert.Equal(t, 1, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{-2147483648}
		res := findPeakElement(nums)
		assert.Equal(t, 0, res)
	})

}
