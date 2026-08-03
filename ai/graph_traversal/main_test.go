package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(2, 1)
		assert.True(t, g.HasCycle())
	})
}
