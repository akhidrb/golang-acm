package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 5
		res := searchInsert(nums, target)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 2
		res := searchInsert(nums, target)
		assert.Equal(t, 1, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 7
		res := searchInsert(nums, target)
		assert.Equal(t, 4, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{1}
		target := 2
		res := searchInsert(nums, target)
		assert.Equal(t, 1, res)
	})

}
