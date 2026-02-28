package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{-1, 2, 1, -4}
		target := 1
		res := threeSumClosest(nums, target)
		exp := 2
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 0, 0}
		target := 1
		res := threeSumClosest(nums, target)
		exp := 0
		assert.Equal(t, exp, res)
	})

}
