package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		res := maxArea(height)
		assert.Equal(t, 49, res)
	})

	t.Run("2", func(t *testing.T) {
		height := []int{1, 1}
		res := maxArea(height)
		assert.Equal(t, 1, res)
	})

}
