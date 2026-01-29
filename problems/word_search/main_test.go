package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		res := exist(board, "ABCCED")
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		res := exist(board, "SEE")
		assert.Equal(t, true, res)
	})

	t.Run("3", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		res := exist(board, "ABCB")
		assert.Equal(t, false, res)
	})

	t.Run("4", func(t *testing.T) {
		board := [][]byte{{'A', 'B'}}
		res := exist(board, "BA")
		assert.Equal(t, true, res)
	})

	t.Run("5", func(t *testing.T) {
		board := [][]byte{{'A', 'B'}, {'C', 'D'}}
		res := exist(board, "BACD")
		assert.Equal(t, true, res)
	})

	t.Run("6", func(t *testing.T) {
		board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
		res := exist(board, "ABCESEEEFS")
		assert.Equal(t, true, res)
	})
}
