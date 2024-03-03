package main

import "testing"

func Test_PrintAorB(t *testing.T) {
	PrintAorB(A{Field1: 1})
	PrintAorB(B{Field2: "2"})
	// PrintAorB(2) // -> type error
}
