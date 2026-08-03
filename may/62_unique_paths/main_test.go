package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		m, n := 3, 2
		exp := 3
		res := uniquePaths(m, n)
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		m, n := 3, 7
		exp := 28
		res := uniquePaths(m, n)
		assert.Equal(t, exp, res)
	})

}
