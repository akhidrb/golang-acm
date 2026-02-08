package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 2}
		res := subsetsWithDup(nums)
		exp := [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 2},
			{2},
			{2, 2},
		}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0}
		res := subsetsWithDup(nums)
		exp := [][]int{
			{},
			{0},
		}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{4, 4, 4, 1, 4}
		res := subsetsWithDup(nums)
		exp := [][]int{
			{},
			{1},
			{1, 4},
			{1, 4, 4},
			{1, 4, 4, 4},
			{1, 4, 4, 4, 4},
			{4},
			{4, 4},
			{4, 4, 4},
			{4, 4, 4, 4},
		}
		assert.ElementsMatch(t, exp, res)
	})

}
