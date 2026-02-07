package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 3
		res := searchMatrix(nums, target)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 13
		res := searchMatrix(nums, target)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := [][]int{{1}}
		target := 2
		res := searchMatrix(nums, target)
		assert.Equal(t, false, res)
	})

}
