package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LongestConsecutiveSequence(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{100, 4, 200, 1, 3, 2}
		val := longestConsecutive(nums)
		assert.Equal(t, 4, val)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		val := longestConsecutive(nums)
		assert.Equal(t, 9, val)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{0, -1}
		val := longestConsecutive(nums)
		assert.Equal(t, 2, val)
	})
	t.Run("3", func(t *testing.T) {
		nums := []int{0, 0, -1}
		val := longestConsecutive(nums)
		assert.Equal(t, 2, val)
	})
}
