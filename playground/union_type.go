package main

import "fmt"

type A struct {
	Field1 int
}

type B struct {
	Field2 string
}

type AorB interface {
	A | B
}

func PrintAorB[T AorB](s T) {
	switch v := any(s).(type) {
	case A:
		fmt.Println(v.Field1)
	case B:
		fmt.Println(v.Field2)
	}
}
