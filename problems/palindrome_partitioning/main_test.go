package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "aab"
		res := partition(s)
		exp := [][]string{{"a", "a", "b"}, {"aa", "b"}}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "a"
		res := partition(s)
		exp := [][]string{{"a"}}
		assert.Equal(t, exp, res)
	})

}
