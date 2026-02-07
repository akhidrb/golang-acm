package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
		res := minimumTotal(triangle)
		assert.Equal(t, 11, res)
	})

	t.Run("2", func(t *testing.T) {
		triangle := [][]int{{-10}}
		res := minimumTotal(triangle)
		assert.Equal(t, -10, res)
	})

}
