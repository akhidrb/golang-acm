package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := fib(10)
		assert.Equal(t, 55, res)
	})

	t.Run("2", func(t *testing.T) {
		res := fibMemo(1000)
		assert.Equal(t, 817770325994397771, res)
	})

}
