package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 5, 0, 3, 5}
		res := minimumOperations(nums)
		assert.Equal(t, 3, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0}
		res := minimumOperations(nums)
		assert.Equal(t, 0, res)
	})

}
