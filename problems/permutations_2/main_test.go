package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 1, 2}
		res := permuteUnique(nums)
		exp := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		res := permuteUnique(nums)
		exp := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}
		assert.ElementsMatch(t, exp, res)
	})
}
