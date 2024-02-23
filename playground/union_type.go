package main

import "fmt"

type Fallback1 struct {
	Field1 int
}

type Fallback2 struct {
	Field2 string
}

type fallback interface {
	Fallback1 | Fallback2
}

func PrintFallback[T fallback](s T) {
	switch v := any(s).(type) {
	case Fallback1:
		fmt.Println(v.Field1)
	case Fallback2:
		fmt.Println(v.Field2)
	}
}
