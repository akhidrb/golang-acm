package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
		res := minPathSum(grid)
		assert.Equal(t, 7, res)
	})

	t.Run("2", func(t *testing.T) {
		grid := [][]int{{1, 2, 3}, {4, 5, 6}}
		res := minPathSum(grid)
		assert.Equal(t, 12, res)
	})

}
