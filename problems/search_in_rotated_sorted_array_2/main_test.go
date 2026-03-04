package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 0
		res := search(nums, target)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 3
		res := search(nums, target)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1, 0, 1, 1, 1}
		target := 0
		res := search(nums, target)
		assert.Equal(t, true, res)
	})

}
