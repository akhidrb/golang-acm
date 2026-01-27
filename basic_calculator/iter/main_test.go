package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "1 + 1"
		res := calculate(s)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "2-1 + 2 "
		res := calculate(s)
		assert.Equal(t, 3, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "(1+(4+5+2)-3)+(6+8)"
		res := calculate(s)
		assert.Equal(t, 23, res)
	})

	t.Run("4", func(t *testing.T) {
		s := "0"
		res := calculate(s)
		assert.Equal(t, 0, res)
	})

	t.Run("5", func(t *testing.T) {
		s := "1-(     -2)"
		res := calculate(s)
		assert.Equal(t, 3, res)
	})

	t.Run("6", func(t *testing.T) {
		s := "-2+ 1"
		res := calculate(s)
		assert.Equal(t, -1, res)
	})

	t.Run("7", func(t *testing.T) {
		s := "2147483647"
		res := calculate(s)
		assert.Equal(t, 2147483647, res)
	})

	t.Run("8", func(t *testing.T) {
		s := "- (3 + (4 + 5))"
		res := calculate(s)
		assert.Equal(t, -12, res)
	})

	t.Run("9", func(t *testing.T) {
		s := "(7)-(0)+(4)"
		res := calculate(s)
		assert.Equal(t, 11, res)
	})

	t.Run("10", func(t *testing.T) {
		s := "(6)-(8)-(7)+(1+(6))"
		res := calculate(s)
		assert.Equal(t, -2, res)
	})
}
