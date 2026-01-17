package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilMap(t *testing.T) {
	var m map[string]int
	assert.Nil(t, m)
	assert.Equal(t, len(m), 0)
	assert.Equal(t, m["hoge"], 0)
}
