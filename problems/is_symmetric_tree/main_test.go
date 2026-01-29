package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p := []any{1, 2, 2, 3, 4, 4, 3}
		pTree := buildTree(p)
		isSame := isSymmetric(pTree)
		assert.Equal(t, true, isSame)
	})

	t.Run("2", func(t *testing.T) {
		p := []any{1, 2, 2, nil, 3, nil, 3}
		pTree := buildTree(p)
		isSame := isSymmetric(pTree)
		assert.Equal(t, false, isSame)
	})

}
