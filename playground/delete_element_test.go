package main

import (
	"fmt"
	"testing"
)

func TestDeleteElement(t *testing.T) {
	n := 2

	// one allocation
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:n], a[n+1:]...)
	fmt.Println(a)

	// no allocation
	b := []int{1, 2, 3, 4, 5}
	b = b[:n+copy(b[n:], b[n+1:])]
	fmt.Println(b)
}
