package main

import (
	"fmt"
	"strings"
	"time"
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
