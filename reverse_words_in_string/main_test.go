package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TextJustification(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "the sky is blue"
		res := reverseWords(s)
		assert.Equal(t, "blue is sky the", res)
	})

	t.Run("2", func(t *testing.T) {
		s := "  hello world  "
		res := reverseWords(s)
		assert.Equal(t, "world hello", res)
	})

	t.Run("3", func(t *testing.T) {
		s := "a good   example"
		res := reverseWords(s)
		assert.Equal(t, "example good a", res)
	})

}
