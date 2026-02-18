package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		res := search(nums, target)
		assert.Equal(t, 4, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 3
		res := search(nums, target)
		assert.Equal(t, -1, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1}
		target := 0
		res := search(nums, target)
		assert.Equal(t, -1, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
		target := 7
		res := search(nums, target)
		assert.Equal(t, 0, res)
	})

	t.Run("5", func(t *testing.T) {
		nums := []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
		target := 3
		res := search(nums, target)
		assert.Equal(t, 6, res)
	})

	t.Run("6", func(t *testing.T) {
		nums := []int{3, 1}
		target := 1
		res := search(nums, target)
		assert.Equal(t, 1, res)
	})

}
