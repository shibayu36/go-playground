package main

import "testing"

func Test_PrintFallback(t *testing.T) {
	PrintFallback(Fallback1{Field1: 1})
	PrintFallback(Fallback2{Field2: "2"})
	// PrintFallback(2) -> type error
}
