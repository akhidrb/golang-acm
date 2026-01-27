package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		str := "dfa12321afd"
		res := secondHighest(str)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		str := "abc1111"
		res := secondHighest(str)
		assert.Equal(t, -1, res)
	})

	t.Run("3", func(t *testing.T) {
		str := "ck077"
		res := secondHighest(str)
		assert.Equal(t, 0, res)
	})

}
