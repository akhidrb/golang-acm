package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		a := NewList([]*NestedInteger{
			NewInteger(1),
			NewInteger(1),
		})
		b := NewInteger(2)
		c := NewList([]*NestedInteger{
			NewInteger(1),
			NewInteger(1),
		})
		root := []*NestedInteger{a, b, c}
		res := depthSum(root)
		assert.Equal(t, 10, res)
	})

}
