package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "I speak Goat Latin"
		res := toGoatLatin(s)
		exp := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "The quick brown fox jumped over the lazy dog"
		res := toGoatLatin(s)
		exp := "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
		assert.Equal(t, exp, res)
	})
}
