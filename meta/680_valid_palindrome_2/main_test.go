package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "aba"
		valid := validPalindrome(s)
		assert.Equal(t, true, valid)
	})

	t.Run("2", func(t *testing.T) {
		s := "abca"
		valid := validPalindrome(s)
		assert.Equal(t, true, valid)
	})

	t.Run("3", func(t *testing.T) {
		s := "abc"
		valid := validPalindrome(s)
		assert.Equal(t, false, valid)
	})

	t.Run("4", func(t *testing.T) {
		s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
		valid := validPalindrome(s)
		assert.Equal(t, true, valid)
	})

}
