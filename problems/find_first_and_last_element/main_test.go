package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		res := searchRange(nums, target)
		exp := []int{3, 4}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 6
		res := searchRange(nums, target)
		exp := []int{-1, -1}
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{}
		target := 0
		res := searchRange(nums, target)
		exp := []int{-1, -1}
		assert.Equal(t, exp, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{1}
		target := 1
		res := searchRange(nums, target)
		exp := []int{0, 0}
		assert.Equal(t, exp, res)
	})

	t.Run("5", func(t *testing.T) {
		nums := []int{1, 4}
		target := 4
		res := searchRange(nums, target)
		exp := []int{1, 1}
		assert.Equal(t, exp, res)
	})

}
