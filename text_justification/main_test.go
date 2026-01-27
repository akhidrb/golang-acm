package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TextJustification(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		words := []string{"This", "is", "an", "example", "of", "text", "justification."}
		res := fullJustify(words, 16)
		expRes := []string{"This    is    an",
			"example  of text",
			"justification.  "}
		assert.Equal(t, expRes, res)
	})

	t.Run("2", func(t *testing.T) {
		words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
		res := fullJustify(words, 16)
		expRes := []string{"What   must   be",
			"acknowledgment  ",
			"shall be        "}
		assert.Equal(t, expRes, res)
	})

	t.Run("3", func(t *testing.T) {
		words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
		res := fullJustify(words, 20)
		expRes := []string{"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  "}
		assert.Equal(t, expRes, res)
	})

}
