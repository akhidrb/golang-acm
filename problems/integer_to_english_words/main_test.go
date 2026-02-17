package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		num := 123
		res := numberToWords(num)
		exp := "One Hundred Twenty Three"
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		num := 12345
		res := numberToWords(num)
		exp := "Twelve Thousand Three Hundred Forty Five"
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		num := 1234567
		res := numberToWords(num)
		exp := "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
		assert.Equal(t, exp, res)
	})

}
