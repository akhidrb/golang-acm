package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
		output := orangesRotting(p)
		assert.Equal(t, 4, output)
	})

	t.Run("2", func(t *testing.T) {
		p := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
		output := orangesRotting(p)
		assert.Equal(t, -1, output)
	})

	t.Run("3", func(t *testing.T) {
		p := [][]int{{0, 2}}
		output := orangesRotting(p)
		assert.Equal(t, 0, output)
	})

}
