package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		num := 19
		res := isHappy(num)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		num := 2
		res := isHappy(num)
		assert.Equal(t, false, res)
	})

}
