package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{3, 2, 1, 5, 6, 4}
		valid := findKthLargest(nums, 2)
		assert.Equal(t, 5, valid)
	})

}
