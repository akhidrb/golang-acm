package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		res := subsets(nums)
		exp := [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0}
		res := subsets(nums)
		exp := [][]int{
			{},
			{0},
		}
		assert.Equal(t, exp, res)
	})

}
