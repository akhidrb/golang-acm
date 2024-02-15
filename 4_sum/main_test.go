package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 0, -1, 0, -2, 2}
		target := 0
		res := fourSum(nums, target)
		assert.Equal(t, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2}
		target := 8
		res := fourSum(nums, target)
		assert.Equal(t, [][]int{{2, 2, 2, 2}}, res)
	})

}
