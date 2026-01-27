package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		roads := [][]int{{0, 1}, {0, 2}, {0, 3}}
		output := minimumFuelCost(roads, 5)
		assert.Equal(t, int64(3), output)
	})

}
