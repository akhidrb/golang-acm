package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := totalNQueens(4)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		res := totalNQueens(1)
		assert.Equal(t, 1, res)
	})
}
