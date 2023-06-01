package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Log struct {
	Host    string
	User    string
	Epoch   int
	Req     string
	Status  int
	Size    int
	Referer string
}

func (l Log) Method() string {
	return strings.Split(l.Req, " ")[0]
}

func (l Log) Path() string {
	return strings.Split(l.Req, " ")[1]
}

func (l Log) Protocol() string {
	return strings.Split(l.Req, " ")[2]
}

func (l Log) Uri() string {
	return fmt.Sprintf("http://%s%s", l.Host, l.Path())
}

func (l Log) Time() string {
	t := time.Unix(int64(l.Epoch), 0).UTC()
	return t.Format("2006-01-02T15:04:05")
}

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
