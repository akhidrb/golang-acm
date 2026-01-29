package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TextJustification(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "abcabcbb"
		res := lengthOfLongestSubstring(str)
		assert.Equal(t, 3, res)
	})

	t.Run("2", func(t *testing.T) {
		str := "bbbbb"
		res := lengthOfLongestSubstring(str)
		assert.Equal(t, 1, res)
	})

	t.Run("3", func(t *testing.T) {
		str := "pwwkew"
		res := lengthOfLongestSubstring(str)
		assert.Equal(t, 3, res)
	})

	t.Run("4", func(t *testing.T) {
		str := "dvdf"
		res := lengthOfLongestSubstring(str)
		assert.Equal(t, 3, res)
	})

}
