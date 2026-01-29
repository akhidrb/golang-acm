package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		x := 121
		isValid := isPalindrome(x)
		assert.Equal(t, true, isValid)
	})

	t.Run("2", func(t *testing.T) {
		x := -121
		isValid := isPalindrome(x)
		assert.Equal(t, false, isValid)
	})

	t.Run("3", func(t *testing.T) {
		x := 10
		isValid := isPalindrome(x)
		assert.Equal(t, false, isValid)
	})

}
