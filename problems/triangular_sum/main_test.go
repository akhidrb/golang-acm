package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		res := triangularSum(nums)
		assert.Equal(t, 8, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{5}
		res := triangularSum(nums)
		assert.Equal(t, 5, res)
	})

}
