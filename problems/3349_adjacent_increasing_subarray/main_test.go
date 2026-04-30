package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
		k := 3
		res := hasIncreasingSubarrays(nums, k)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
		k := 5
		res := hasIncreasingSubarrays(nums, k)
		assert.Equal(t, false, res)
	})
}
