package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "III"
		num := romanToInt(s)
		assert.Equal(t, 3, num)
	})

	t.Run("2", func(t *testing.T) {
		s := "LVIII"
		num := romanToInt(s)
		assert.Equal(t, 58, num)
	})

	t.Run("2", func(t *testing.T) {
		s := "MCMXCIV"
		num := romanToInt(s)
		assert.Equal(t, 1994, num)
	})

}
