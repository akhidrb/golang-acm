package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "bcabc"
		res := removeDuplicateLetters(s)
		exp := "abc"
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "cbacdcbc"
		res := removeDuplicateLetters(s)
		exp := "acdb"
		assert.Equal(t, exp, res)
	})

}
