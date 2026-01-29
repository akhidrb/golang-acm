package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := generate(5)
		exp := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		res := generate(1)
		exp := [][]int{{1}}
		assert.Equal(t, exp, res)
	})

}
