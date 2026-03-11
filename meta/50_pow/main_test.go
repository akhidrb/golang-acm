package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		res := myPow(2.0, 10)
		assert.Equal(t, 1024.0, res)
	})

	t.Run("2", func(t *testing.T) {
		res := myPow(2.1, 3)
		res = math.Round(res*1000) / 1000
		assert.Equal(t, 9.261, res)
	})

	t.Run("3", func(t *testing.T) {
		res := myPow(2.0, -2)
		assert.Equal(t, 0.25, res)
	})

}
