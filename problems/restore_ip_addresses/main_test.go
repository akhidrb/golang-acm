package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "25525511135"
		res := restoreIpAddresses(s)
		exp := []string{"255.255.11.135", "255.255.111.35"}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "0000"
		res := restoreIpAddresses(s)
		exp := []string{"0.0.0.0"}
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "101023"
		res := restoreIpAddresses(s)
		exp := []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}
		assert.Equal(t, exp, res)
	})

}
