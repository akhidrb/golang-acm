package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		candidates := []int{10, 1, 2, 7, 6, 1, 5}
		target := 8
		res := combinationSum2(candidates, target)
		exp := [][]int{
			{1, 1, 6},
			{1, 2, 5},
			{1, 7},
			{2, 6},
		}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		candidates := []int{2, 5, 2, 1, 2}
		target := 5
		res := combinationSum2(candidates, target)
		exp := [][]int{
			{1, 2, 2},
			{5},
		}
		assert.ElementsMatch(t, exp, res)
	})
}
