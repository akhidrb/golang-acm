package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		grid := [][]int{
			{1, 0},
			{0, 1},
		}
		res := largestIsland(grid)
		assert.Equal(t, 3, res)
	})

	t.Run("2", func(t *testing.T) {
		grid := [][]int{
			{1, 1},
			{1, 0},
		}
		res := largestIsland(grid)
		assert.Equal(t, 4, res)
	})

	t.Run("3", func(t *testing.T) {
		grid := [][]int{
			{1, 1},
			{1, 1},
		}
		res := largestIsland(grid)
		assert.Equal(t, 4, res)
	})
}
