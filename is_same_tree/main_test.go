package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p, q := []any{1, 2, 3}, []any{1, 2, 3}
		pTree, qTree := buildTree(p), buildTree(q)
		isSame := isSameTree(pTree, qTree)
		assert.Equal(t, true, isSame)
	})

	t.Run("2", func(t *testing.T) {
		p, q := []any{1, 2}, []any{1, nil, 2}
		pTree, qTree := buildTree(p), buildTree(q)
		isSame := isSameTree(pTree, qTree)
		assert.Equal(t, false, isSame)
	})

}
