package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
		res := uniquePathsWithObstacles(obstacleGrid)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		obstacleGrid := [][]int{{0, 1}, {0, 0}}
		res := uniquePathsWithObstacles(obstacleGrid)
		assert.Equal(t, 1, res)
	})

}
