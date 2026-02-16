package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "aabb" // abba, baab
		res := minMovesToMakePalindrome(s)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "letelt" // lettel , lteetl, tellet, tleelt, etllte, elttle
		res := minMovesToMakePalindrome(s)
		assert.Equal(t, 2, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "eqvvhtcsaaqtqesvvqch" //
		res := minMovesToMakePalindrome(s)
		assert.Equal(t, 17, res)
	})
}
