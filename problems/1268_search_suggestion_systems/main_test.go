package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
		searchWord := "mouse"
		res := suggestedProducts(products, searchWord)
		exp := [][]string{
			{"mobile", "moneypot", "monitor"},
			{"mobile", "moneypot", "monitor"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"}}
		assert.Equal(t, exp, res)
	})

}
