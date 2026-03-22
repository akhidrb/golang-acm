package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		nextPermutation(nums)
		exp := []int{1, 3, 2}
		assert.Equal(t, exp, nums)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{3, 2, 1}
		nextPermutation(nums)
		exp := []int{1, 2, 3}
		assert.Equal(t, exp, nums)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1, 1, 5}
		nextPermutation(nums)
		exp := []int{1, 5, 1}
		assert.Equal(t, exp, nums)
	})

}
