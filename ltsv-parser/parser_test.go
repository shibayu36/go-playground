package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	p := Parser{filename: "./testdata/log.ltsv"}
	logs, err := p.Parse()

	assert.Nil(t, err)

	assert.Equal(t, 3, len(logs))

	assert.Equal(t, 1372694390, logs[0].Epoch)
	assert.Equal(t, 1372794390, logs[1].Epoch)
	assert.Equal(t, 1372894390, logs[2].Epoch)
}

func TestLineToLog(t *testing.T) {
	t.Run("valid line", func(t *testing.T) {
		p := Parser{filename: "./testdata/log.ltsv"}
		log, err := p.lineToLog("host:127.0.0.1	user:-	epoch:1372894390	req:GET /apache_pb.gif HTTP/1.0	status:302	size:9999	referer:http://www.example.com/start.html")

		assert.Nil(t, err)

		assert.Equal(t, "127.0.0.1", log.Host)
		assert.Equal(t, "", log.User)
		assert.Equal(t, 1372894390, log.Epoch)
		assert.Equal(t, "GET /apache_pb.gif HTTP/1.0", log.Req)
		assert.Equal(t, 302, log.Status)
		assert.Equal(t, 9999, log.Size)
		assert.Equal(t, "http://www.example.com/start.html", log.Referer)
	})

	t.Run("invalid line", func(t *testing.T) {
		p := Parser{filename: "./testdata/log.ltsv"}
		log, err := p.lineToLog("host:127.0.0.1	user:-	epoch:aiu	req:GET /apache_pb.gif HTTP/1.0	status:302	size:9999	referer:http://www.example.com/start.html")

		assert.NotNil(t, err)
		assert.Nil(t, log)
	})
}
