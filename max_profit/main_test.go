package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxProfit(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{7, 1, 5, 3, 6, 4}
		res := maxProfit(nums)
		assert.Equal(t, 7, res)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		res := maxProfit(nums)
		assert.Equal(t, 4, res)
	})
	t.Run("3", func(t *testing.T) {
		nums := []int{7, 6, 4, 3, 1}
		res := maxProfit(nums)
		assert.Equal(t, 0, res)
	})
	t.Run("4", func(t *testing.T) {
		nums := []int{1, 2, 5, 4, 7}
		res := maxProfit(nums)
		assert.Equal(t, 7, res)
	})
}
