package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int64{6, 3, 5, 1, 12}
		res := riddle(nums)
		assert.Equal(t, []int64{12, 3, 3, 1, 1}, res)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int64{2, 6, 1, 12}
		res := riddle(nums)
		assert.Equal(t, []int64{12, 2, 1, 1}, res)
	})
}
