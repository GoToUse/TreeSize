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

func TestMultiExists(t *testing.T) {
	is := assert.New(t)

	// case 1
	preBool, expect := []bool{true, false, false}, 1

	res := multiExists(preBool, expect)
	is.Equal(false, res)

	// case 2
	preBool1, expect1 := []bool{true, true, false}, 1

	res1 := multiExists(preBool1, expect1)
	is.Equal(true, res1)

	// case 3
	preBool2, expect2 := []bool{true, true, true}, 1

	res2 := multiExists(preBool2, expect2)
	is.Equal(true, res2)
}
