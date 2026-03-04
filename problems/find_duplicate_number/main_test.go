package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 3, 4, 2, 2}
		res := findDuplicate(nums)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{3, 1, 3, 4, 2}
		res := findDuplicate(nums)
		assert.Equal(t, 3, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{3, 3, 3, 3, 3}
		res := findDuplicate(nums)
		assert.Equal(t, 3, res)
	})

}
