package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := isMatch("aa", "a")
		assert.Equal(t, false, res)
	})

	t.Run("2", func(t *testing.T) {
		res := isMatch("aa", "a*")
		assert.Equal(t, true, res)
	})

	t.Run("3", func(t *testing.T) {
		res := isMatch("ab", ".*")
		assert.Equal(t, true, res)
	})

	t.Run("4", func(t *testing.T) {
		res := isMatch("aab", "c*a*b")
		assert.Equal(t, true, res)
	})

}
