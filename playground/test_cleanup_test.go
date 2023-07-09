package main

import (
	"fmt"
	"testing"
)

func TestCleanup(t *testing.T) {

	t.Cleanup(func() {
		fmt.Println("cleanup: root")
	})
	fmt.Println("start: root")

	t.Run("subtest1", func(t *testing.T) {
		t.Cleanup(func() {
			fmt.Println("cleanup: subtest1")
		})
		fmt.Println("start: subtest1")

		t.Run("subtest1-1", func(t *testing.T) {
			t.Cleanup(func() {
				fmt.Println("cleanup: subtest1-1")
			})
			fmt.Println("start: subtest1-1")
		})

		t.Run("subtest1-2", func(t *testing.T) {
			t.Cleanup(func() {
				fmt.Println("cleanup: subtest1-2")
			})
			fmt.Println("start: subtest1-2")
		})
	})

	t.Run("subtest2", func(t *testing.T) {
		t.Cleanup(func() {
			fmt.Println("cleanup: subtest2")
		})
		fmt.Println("start: subtest2")

		t.Run("subtest2-1", func(t *testing.T) {
			t.Cleanup(func() {
				fmt.Println("cleanup: subtest2-1")
			})
			fmt.Println("start: subtest2-1")
		})
	})
}
