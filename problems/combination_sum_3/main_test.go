package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		k := 3
		n := 7
		res := combinationSum3(k, n)
		exp := [][]int{{1, 2, 4}}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		k := 3
		n := 9
		res := combinationSum3(k, n)
		exp := [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}
		assert.Equal(t, exp, res)
	})

}
