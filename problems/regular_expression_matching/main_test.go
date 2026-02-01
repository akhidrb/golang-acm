package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "aa"
		p := "a"
		res := isMatch(s, p)
		assert.Equal(t, false, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "aa"
		p := "a*"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "ab"
		p := ".*"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("4", func(t *testing.T) {
		s := "aaaaaabbbbbbbbbcdeffff"
		p := "a*b*cdef*"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("5", func(t *testing.T) {
		s := "bbbbbbbbbcdeffff"
		p := "a*b*cdef*"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("6", func(t *testing.T) {
		s := "aaaaacdeffff"
		p := "a*b*cdef*"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("7", func(t *testing.T) {
		s := "mississippi"
		p := "mis*is*ip*."
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("8", func(t *testing.T) {
		s := "ab"
		p := ".*c"
		res := isMatch(s, p)
		assert.Equal(t, false, res)
	})

	t.Run("9", func(t *testing.T) {
		s := "aaa"
		p := "a*a"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

	t.Run("10", func(t *testing.T) {
		s := "aaa"
		p := "a*aaa"
		res := isMatch(s, p)
		assert.Equal(t, true, res)
	})

}
