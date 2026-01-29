package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "barfoothefoobarman"
		words := []string{"foo", "bar"}
		res := findSubstring(str, words)
		assert.Equal(t, []int{0, 9}, res)
	})

	t.Run("2", func(t *testing.T) {
		str := "wordgoodgoodgoodbestword"
		words := []string{"word", "good", "best", "word"}
		res := findSubstring(str, words)
		assert.Equal(t, []int{}, res)
	})

	t.Run("3", func(t *testing.T) {
		str := "barfoofoobarthefoobarman"
		words := []string{"bar", "foo", "the"}
		res := findSubstring(str, words)
		assert.Equal(t, []int{6, 9, 12}, res)
	})

	t.Run("4", func(t *testing.T) {
		str := "wordgoodgoodgoodbestword"
		words := []string{"word", "good", "best", "good"}
		res := findSubstring(str, words)
		assert.Equal(t, []int{8}, res)
	})

	t.Run("5", func(t *testing.T) {
		str := "aaaaaaaaaaaaaa"
		words := []string{"aa", "aa"}
		res := findSubstring(str, words)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, res)
	})

}
