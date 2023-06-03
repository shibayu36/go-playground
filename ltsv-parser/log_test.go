package main

import (
	"testing"
)

var log = Log{
	Host:    "testhost",
	User:    "testuser",
	Epoch:   1685776955,
	Req:     "GET /path HTTP/1.1",
	Status:  200,
	Size:    100,
	Referer: "http://example.com",
}

func TestMethod(t *testing.T) {
	if log.Method() != "GET" {
		t.Errorf("Method() = %s; want GET", log.Method())
	}
}

func TestPath(t *testing.T) {
	if log.Path() != "/path" {
		t.Errorf("Path() = %s; want /path", log.Path())
	}
}

func TestProtocol(t *testing.T) {
	if log.Protocol() != "HTTP/1.1" {
		t.Errorf("Protocol() = %s; want HTTP/1.1", log.Protocol())
	}
}

func TestUri(t *testing.T) {
	if log.Uri() != "http://testhost/path" {
		t.Errorf("Uri() = %s; want http://testhost/path", log.Uri())
	}
}

func TestTime(t *testing.T) {
	expectedTime := "2023-06-03T07:22:35"
	if log.Time() != expectedTime {
		t.Errorf("Time() = %s; want %s", log.Time(), expectedTime)
	}
}
