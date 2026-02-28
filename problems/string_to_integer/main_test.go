package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "42"
		res := myAtoi(s)
		exp := 42
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		s := " -042"
		res := myAtoi(s)
		exp := -42
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "1337c0d3"
		res := myAtoi(s)
		exp := 1337
		assert.Equal(t, exp, res)
	})

	t.Run("4", func(t *testing.T) {
		s := "0-1"
		res := myAtoi(s)
		exp := 0
		assert.Equal(t, exp, res)
	})

	t.Run("5", func(t *testing.T) {
		s := "words and 987"
		res := myAtoi(s)
		exp := 0
		assert.Equal(t, exp, res)
	})

	t.Run("6", func(t *testing.T) {
		s := "-91283472332"
		res := myAtoi(s)
		exp := -2147483648
		assert.Equal(t, exp, res)
	})

	t.Run("7", func(t *testing.T) {
		s := "9223372036854775808"
		res := myAtoi(s)
		exp := 2147483647
		assert.Equal(t, exp, res)
	})

}
