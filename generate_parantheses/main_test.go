package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := generateParenthesis(3)
		exp := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		res := generateParenthesis(1)
		exp := []string{"()"}
		assert.Equal(t, exp, res)
	})
}
