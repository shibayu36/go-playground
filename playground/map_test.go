package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	hoge := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println(hoge["foo"])
	fmt.Println(hoge["bar"])
	fmt.Println(hoge["baz"])
}
