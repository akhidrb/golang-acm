package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		res := rob(nums)
		assert.Equal(t, 4, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{2, 7, 9, 3, 1}
		res := rob(nums)
		assert.Equal(t, 12, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{100, 7, 9, 200, 1}
		res := rob(nums)
		assert.Equal(t, 300, res)
	})

}
