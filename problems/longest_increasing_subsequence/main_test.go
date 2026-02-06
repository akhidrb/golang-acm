package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		res := lengthOfLIS(nums)
		assert.Equal(t, 4, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 2, 3}
		res := lengthOfLIS(nums)
		assert.Equal(t, 4, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7, 7, 7}
		res := lengthOfLIS(nums)
		assert.Equal(t, 1, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
		res := lengthOfLIS(nums)
		assert.Equal(t, 6, res)
	})

}
