package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IntegerToRoman(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := intToRoman(3)
		assert.Equal(t, "III", res)
	})
	t.Run("2", func(t *testing.T) {
		res := intToRoman(58)
		assert.Equal(t, "LVIII", res)
	})
	t.Run("3", func(t *testing.T) {
		res := intToRoman(1994)
		assert.Equal(t, "MCMXCIV", res)
	})
	t.Run("4", func(t *testing.T) {
		res := intToRoman(10)
		assert.Equal(t, "X", res)
	})
	t.Run("5", func(t *testing.T) {
		res := intToRoman(100)
		assert.Equal(t, "C", res)
	})
	t.Run("6", func(t *testing.T) {
		res := intToRoman(140)
		assert.Equal(t, "CXL", res)
	})
	t.Run("7", func(t *testing.T) {
		res := intToRoman(2000)
		assert.Equal(t, "MM", res)
	})
}
