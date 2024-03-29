package main

import (
	"testing"

	"github.com/GoToUse/treeprint"
	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	is := assert.New(t)
	tree := treeprint.New()

	size, err := Parallel(".", tree)

	is.Nil(err)

	is.GreaterOrEqual(size, int64(0))
}
