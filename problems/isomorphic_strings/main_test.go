package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s1 := "egg"
		s2 := "add"
		res := isIsomorphic(s1, s2)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		s1 := "f11"
		s2 := "b23"
		res := isIsomorphic(s1, s2)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		s1 := "paper"
		s2 := "title"
		res := isIsomorphic(s1, s2)
		assert.Equal(t, true, res)
	})

	t.Run("4", func(t *testing.T) {
		s1 := "badc"
		s2 := "baba"
		res := isIsomorphic(s1, s2)
		assert.Equal(t, false, res)
	})

}
