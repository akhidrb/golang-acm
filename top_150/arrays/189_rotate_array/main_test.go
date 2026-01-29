package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		rotate(nums, k)
		exp := []int{5, 6, 7, 1, 2, 3, 4}
		assert.Equal(t, exp, nums)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{-1, -100, 3, 99}
		k := 2
		rotate(nums, k)
		exp := []int{3, 99, -1, -100}
		assert.Equal(t, exp, nums)
	})

}
