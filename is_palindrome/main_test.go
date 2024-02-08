package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TextJustification(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "A man, a plan, a canal: Panama"
		res := isPalindrome(str)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		str := "race a car"
		res := isPalindrome(str)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		str := " "
		res := isPalindrome(str)
		assert.Equal(t, true, res)
	})

}
