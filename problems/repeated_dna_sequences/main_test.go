package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
		res := findRepeatedDnaSequences(str)
		exp := []string{"AAAAACCCCC", "CCCCCAAAAA"}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		str := "AAAAAAAAAAAAA"
		res := findRepeatedDnaSequences(str)
		exp := []string{"AAAAAAAAAA"}
		assert.Equal(t, exp, res)
	})

}
