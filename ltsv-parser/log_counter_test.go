package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogCounterCountError(t *testing.T) {
	lc := LogCounter{
		Logs: []Log{
			{
				Status: 200,
			},
		},
	}
	assert.Equal(t, 0, lc.CountError())
}
