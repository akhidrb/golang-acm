package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		word := "internationalization"
		abbr := "i12iz4n"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, true, valid)
	})

	t.Run("2", func(t *testing.T) {
		word := "apple"
		abbr := "a2e"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, false, valid)
	})

	t.Run("3", func(t *testing.T) {
		word := "internationalization"
		abbr := "i5a11o1"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, true, valid)
	})

	t.Run("4", func(t *testing.T) {
		word := "abbde"
		abbr := "a1b01e"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, false, valid)
	})

	t.Run("5", func(t *testing.T) {
		word := "substitution"
		abbr := "s55n"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, false, valid)
	})

	t.Run("6", func(t *testing.T) {
		word := "substitution"
		abbr := "s010n"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, false, valid)
	})

	t.Run("7", func(t *testing.T) {
		word := "substitution"
		abbr := "s0ubstitution"
		valid := validWordAbbreviation(word, abbr)
		assert.Equal(t, false, valid)
	})

}
