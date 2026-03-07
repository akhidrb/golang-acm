package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "lee(t(c)o)de)"
		num := minRemoveToMakeValid(s)
		exp := "lee(t(c)o)de"
		assert.Equal(t, exp, num)
	})

	t.Run("2", func(t *testing.T) {
		s := "a)b(c)d"
		num := minRemoveToMakeValid(s)
		exp := "ab(c)d"
		assert.Equal(t, exp, num)
	})

	t.Run("2", func(t *testing.T) {
		s := "))(("
		num := minRemoveToMakeValid(s)
		exp := ""
		assert.Equal(t, exp, num)
	})

}
