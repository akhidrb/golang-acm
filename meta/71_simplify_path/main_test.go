package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "/home/"
		valid := simplifyPath(s)
		exp := "/home"
		assert.Equal(t, exp, valid)
	})

	t.Run("2", func(t *testing.T) {
		s := "/home//foo/"
		valid := simplifyPath(s)
		exp := "/home/foo"
		assert.Equal(t, exp, valid)
	})

	t.Run("3", func(t *testing.T) {
		s := "/home/user/Documents/../Pictures"
		valid := simplifyPath(s)
		exp := "/home/user/Pictures"
		assert.Equal(t, exp, valid)
	})

	t.Run("4", func(t *testing.T) {
		s := "/../"
		valid := simplifyPath(s)
		exp := "/"
		assert.Equal(t, exp, valid)
	})

	t.Run("5", func(t *testing.T) {
		s := "/.../a/../b/c/../d/./"
		valid := simplifyPath(s)
		exp := "/.../b/d"
		assert.Equal(t, exp, valid)
	})

}
