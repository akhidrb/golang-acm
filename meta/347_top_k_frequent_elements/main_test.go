package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		res := topKFrequent(nums, 2)
		assert.Equal(t, []int{1, 2}, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1}
		res := topKFrequent(nums, 1)
		assert.Equal(t, []int{1}, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
		res := topKFrequent(nums, 2)
		assert.Equal(t, []int{1, 2}, res)
	})

}
