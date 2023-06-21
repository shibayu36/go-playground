package main

import (
	"fmt"
	"testing"
)

// func main() {
// 	// if using buffered channel, it will not cause deadlock until 2 elements are added
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

func TestChannel(t *testing.T) {
	// if make(chan int) is used without goroutine, it will cause deadlock
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
