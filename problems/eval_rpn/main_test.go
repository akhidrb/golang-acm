package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		tokens := []string{"2", "1", "+", "3", "*"}
		num := evalRPN(tokens)
		assert.Equal(t, 9, num)
	})

	t.Run("2", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		num := evalRPN(tokens)
		assert.Equal(t, 6, num)
	})

	t.Run("3", func(t *testing.T) {
		tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
		num := evalRPN(tokens)
		assert.Equal(t, 22, num)
	})

}
