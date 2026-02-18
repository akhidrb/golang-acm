package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "Hello World"
		res := lengthOfLastWord(s)
		assert.Equal(t, 5, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "   fly me   to   the moon  "
		res := lengthOfLastWord(s)
		assert.Equal(t, 4, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "luffy is still joyboy"
		res := lengthOfLastWord(s)
		assert.Equal(t, 6, res)
	})

}
