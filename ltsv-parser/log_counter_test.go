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

func TestLogCounterGroupByLog(t *testing.T) {
	parser := Parser{filename: "./testdata/log_for_counter.ltsv"}
	logs, _ := parser.Parse()
	lc := LogCounter{Logs: logs}
	group := lc.GroupByUser()

	assert.Equal(t, 3, len(group))

	assert.Equal(t, 2, len(group["frank"]))
	assert.Equal(t, 1372694390, group["frank"][0].Epoch)
	assert.Equal(t, 1372694391, group["frank"][1].Epoch)

	assert.Equal(t, 1, len(group["john"]))
	assert.Equal(t, 1372794390, group["john"][0].Epoch)

	assert.Equal(t, 1, len(group["guest"]))
	assert.Equal(t, 1372894390, group["guest"][0].Epoch)
}
