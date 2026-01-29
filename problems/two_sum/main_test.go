package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		k := 9
		two := twoSum(nums, k)
		assert.Equal(t, []int{1, 2}, two)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{2, 3, 4}
		k := 6
		two := twoSum(nums, k)
		assert.Equal(t, []int{1, 3}, two)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{-1, 0}
		k := -1
		two := twoSum(nums, k)
		assert.Equal(t, []int{1, 2}, two)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{-20, -10, -3, 10, 25}
		k := 0
		two := twoSum(nums, k)
		assert.Equal(t, []int{2, 4}, two)
	})
}
