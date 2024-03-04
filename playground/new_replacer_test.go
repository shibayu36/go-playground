package main

import (
	"strings"
	"testing"
)

func BenchmarkNewReplacer(b *testing.B) {
	// NewReplacer speed
	for i := 0; i < b.N; i++ {
		strings.NewReplacer(
			"hoge", "foo",
			"moge", "bar",
			"huga", "baz",
		)
	}
}
