package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		haystack, needle := "sadbutsad", "sad"
		res := strStr(haystack, needle)
		assert.Equal(t, 0, res)
	})

	t.Run("2", func(t *testing.T) {
		haystack, needle := "leetcode", "leeto"
		res := strStr(haystack, needle)
		assert.Equal(t, -1, res)
	})

	t.Run("3", func(t *testing.T) {
		haystack, needle := "a", "a"
		res := strStr(haystack, needle)
		assert.Equal(t, 0, res)
	})

}
