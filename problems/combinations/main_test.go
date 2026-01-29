package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		n, k := 4, 2
		res := combine(n, k)
		exp := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		n, k := 1, 1
		res := combine(n, k)
		exp := [][]int{{1}}
		assert.Equal(t, exp, res)
	})
}
