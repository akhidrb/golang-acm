package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		digits := "23"
		res := letterCombinations(digits)
		assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, res)
	})

	t.Run("2", func(t *testing.T) {
		digits := ""
		res := letterCombinations(digits)
		assert.Equal(t, []string{}, res)
	})

	t.Run("3", func(t *testing.T) {
		digits := "2"
		res := letterCombinations(digits)
		assert.Equal(t, []string{"a", "b", "c"}, res)
	})

}
