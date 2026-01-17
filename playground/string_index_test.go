package main

import (
	"fmt"
	"testing"
)

func Test_StringIndex(t *testing.T) {
	str := "Hello, World"
	fmt.Println(string(str[8]))
	fmt.Println(str[:3])
	// fmt.Println(str[100])
}
