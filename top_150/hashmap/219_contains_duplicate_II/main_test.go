package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums, k := []int{1, 2, 3, 1}, 3
		res := containsNearbyDuplicate(nums, k)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		nums, k := []int{1, 0, 1, 1}, 1
		res := containsNearbyDuplicate(nums, k)
		assert.Equal(t, true, res)
	})

	t.Run("3", func(t *testing.T) {
		nums, k := []int{1, 2, 3, 1, 2, 3}, 2
		res := containsNearbyDuplicate(nums, k)
		assert.Equal(t, false, res)
	})

}
