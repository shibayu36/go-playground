package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogCounterCountError(t *testing.T) {
	parser := Parser{filename: "./testdata/log_for_counter.ltsv"}
	logs, _ := parser.Parse()
	lc := LogCounter{
		Logs: logs,
	}
	assert.Equal(t, 2, lc.CountError())
}
