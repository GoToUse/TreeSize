package main

import (
	"testing"

	"github.com/GoToUse/treeprint"
	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	is := assert.New(t)
	tree := treeprint.New()

	size, err := Parallel("/Users/dapeng/Desktop/code/Git", tree)

	is.Nil(err)

	is.Equal(size, int64(2636490533))
}
