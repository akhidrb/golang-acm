package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
		k := 4
		res := findKthLargest(nums, k)
		assert.Equal(t, 4, res)
	})

}
