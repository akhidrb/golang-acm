package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s1 := "anagram"
		s2 := "nagaram"
		res := isAnagram(s1, s2)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		s1 := "rat"
		s2 := "car"
		res := isAnagram(s1, s2)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		s1 := "a"
		s2 := "ab"
		res := isAnagram(s1, s2)
		assert.Equal(t, false, res)
	})
}
