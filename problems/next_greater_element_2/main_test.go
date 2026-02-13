package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 1}
		res := nextGreaterElements(nums)
		exp := []int{2, -1, 2}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 3}
		res := nextGreaterElements(nums)
		exp := []int{2, 3, 4, -1, 4}
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{5, 4, 3, 2, 1}
		res := nextGreaterElements(nums)
		exp := []int{-1, 5, 5, 5, 5}
		assert.Equal(t, exp, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}
		res := nextGreaterElements(nums)
		exp := []int{120, 11, 120, 120, 123, 123, -1, 100, 100, 100}
		assert.Equal(t, exp, res)
	})

	t.Run("5", func(t *testing.T) {
		nums := []int{3, -2, -1}
		res := nextGreaterElements(nums)
		exp := []int{-1, -1, 3}
		assert.Equal(t, exp, res)
	})

}
