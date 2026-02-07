package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "leetcode"
		wordDict := []string{"leet", "code"}
		res := wordBreak(s, wordDict)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "applepenapple"
		wordDict := []string{"apple", "pen"}
		res := wordBreak(s, wordDict)
		assert.Equal(t, true, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "catsandog"
		wordDict := []string{"cats", "dog", "sand", "and", "cat"}
		res := wordBreak(s, wordDict)
		assert.Equal(t, false, res)
	})

	t.Run("4", func(t *testing.T) {
		s := "aaaaaaa"
		wordDict := []string{"aaaa", "aaa"}
		res := wordBreak(s, wordDict)
		assert.Equal(t, true, res)
	})

	t.Run("5", func(t *testing.T) {
		s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
		wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
		res := wordBreak(s, wordDict)
		assert.Equal(t, false, res)
	})

}
