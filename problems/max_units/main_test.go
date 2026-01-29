package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
		truckSize := 4
		res := maximumUnits(boxTypes, truckSize)
		assert.Equal(t, 8, res)
	})

	t.Run("1", func(t *testing.T) {
		boxTypes := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
		truckSize := 10
		res := maximumUnits(boxTypes, truckSize)
		assert.Equal(t, 91, res)
	})

}
