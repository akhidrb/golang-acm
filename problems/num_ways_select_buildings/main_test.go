package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "001101"
		res := numberOfWays(s)
		assert.Equal(t, int64(6), res)
	})

	t.Run("2", func(t *testing.T) {
		s := "11100"
		res := numberOfWays(s)
		assert.Equal(t, int64(0), res)
	})

}
