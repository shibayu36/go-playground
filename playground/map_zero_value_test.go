package main

import (
	"fmt"
	"testing"
)

func TestMapZeroValue(t *testing.T) {
	hoge := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println(hoge["foo"])
	fmt.Println(hoge["bar"])
	fmt.Println(hoge["baz"])

	val, ok := hoge["baz"]
	fmt.Println(val, ok)
}
