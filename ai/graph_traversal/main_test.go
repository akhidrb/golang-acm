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

	t.Run("2", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(2, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 5)
		assert.True(t, g.ReachesEnd(1, 5))
	})
}
