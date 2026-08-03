package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		vendors := []string{"acme", "globex", "acme", "initech", "acme", "globex", "umbrella"}
		k := 2
		flagged := flaggedVendors(vendors, k)
		assert.Equal(t, []string{"acme", "globex"}, flagged)
	})

}
