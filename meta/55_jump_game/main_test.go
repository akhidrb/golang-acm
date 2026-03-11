package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		res := canJump(nums)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}
		res := canJump(nums)
		assert.Equal(t, false, res)
	})
}
