package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		pattern, s := "abba", "dog cat cat dog"
		res := wordPattern(pattern, s)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		pattern, s := "abba", "dog cat cat fish"
		res := wordPattern(pattern, s)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		pattern, s := "aaaa", "dog cat cat dog"
		res := wordPattern(pattern, s)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		pattern, s := "abba", "dog dog dog dog"
		res := wordPattern(pattern, s)
		assert.Equal(t, false, res)
	})

}
