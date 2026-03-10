package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "A man, a plan, a canal: Panama"
		res := isPalindrome(s)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "race a car"
		res := isPalindrome(s)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		s := " "
		res := isPalindrome(s)
		assert.Equal(t, true, res)
	})

}
