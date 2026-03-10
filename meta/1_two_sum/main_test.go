package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		res := twoSum(nums, target)
		exp := []int{0, 1}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		res := twoSum(nums, target)
		exp := []int{1, 2}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		res := twoSum(nums, target)
		exp := []int{0, 1}
		assert.ElementsMatch(t, exp, res)
	})
}
