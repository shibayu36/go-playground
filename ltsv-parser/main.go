package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	log := Log{
		Host:    "127.0.0.1",
		User:    "-",
		Epoch:   1549573860,
		Req:     "GET / HTTP/1.1",
		Status:  200,
		Size:    777,
		Referer: "https://example.com/",
	}
	spew.Dump(log)
	println(log.Method())
	println(log.Path())
	println(log.Protocol())
	println(log.Uri())
	println(log.Time())
}
