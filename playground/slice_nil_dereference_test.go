package main

import (
	"fmt"
	"testing"
)

type HogeInt int

func (h HogeInt) String() string {
	return fmt.Sprintf("%d", h)
}

type HogeSlice []int

func (h HogeSlice) String() string {
	return "hoge"
}

type HogeMap map[string]int

func (h HogeMap) String() string {
	return "fuga"
}

type HogeStruct struct {
}

func (h HogeStruct) String() string {
	return "piyo"
}

func TestSliceNilDereference(t *testing.T) {
	// var hogeInt *HogeInt
	// fmt.Println(hogeInt.String())
	var hogeSlice HogeSlice
	fmt.Println(hogeSlice.String())
	var hogeMap HogeMap
	fmt.Println(hogeMap.String())
	var hogeStruct HogeStruct
	fmt.Println(hogeStruct.String())
}
