package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TextJustification(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "abc"
		str2 := "ahbgdc"
		res := isSubsequence(str, str2)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		str := "axc"
		str2 := "ahbgdc"
		res := isSubsequence(str, str2)
		assert.Equal(t, false, res)
	})

}
