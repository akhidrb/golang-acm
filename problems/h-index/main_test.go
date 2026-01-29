package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HIndex(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{3, 0, 6, 1, 5}
		res := hIndex(nums)
		assert.Equal(t, 3, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1, 3, 1}
		res := hIndex(nums)
		assert.Equal(t, 1, res)
	})
}
