package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "ADOBECODEBANC"
		str2 := "ABC"
		res := minWindow(str, str2)
		assert.Equal(t, "BANC", res)
	})

	t.Run("2", func(t *testing.T) {
		str := "a"
		str2 := "a"
		res := minWindow(str, str2)
		assert.Equal(t, "a", res)
	})

	t.Run("3", func(t *testing.T) {
		str := "a"
		str2 := "aa"
		res := minWindow(str, str2)
		assert.Equal(t, "", res)
	})

}
