package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{3, 1, 2, 4}
		res := sumSubarrayMins(nums)
		assert.Equal(t, 17, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{11, 81, 94, 43, 3}
		res := sumSubarrayMins(nums)
		assert.Equal(t, 444, res)
	})

}
