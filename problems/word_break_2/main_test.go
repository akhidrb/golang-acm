package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "catsanddog"
		wordDict := []string{"cat", "cats", "and", "sand", "dog"}
		res := wordBreak(s, wordDict)
		exp := []string{"cats and dog", "cat sand dog"}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "pineapplepenapple"
		wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
		res := wordBreak(s, wordDict)
		exp := []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}
		assert.ElementsMatch(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "catsandog"
		wordDict := []string{"cats", "dog", "sand", "and", "cat"}
		res := wordBreak(s, wordDict)
		exp := []string{}
		assert.ElementsMatch(t, exp, res)
	})

}
