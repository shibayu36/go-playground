package main

import (
	"errors"
	"fmt"
	"testing"

	coerrors "github.com/cockroachdb/errors"
	pkgerrors "github.com/pkg/errors"
)

func TestCockroachError(t *testing.T) {
	err1 := errorFunc()
	errWrapped1 := pkgerrors.Wrap(err1, "wrapped")
	fmt.Printf("%+v\n", errWrapped1)

	fmt.Printf("\n-------------\n")

	err2 := errorFunc()
	errWrapped2 := coerrors.Wrap(err2, "wrapped")
	fmt.Printf("%+v\n", errWrapped2)
}

func errorWrapFunc(f func(error) error) error {
	err := errorFunc()
	return f(err)
}

func errorFunc() error {
	return errors.New("error1")
}
