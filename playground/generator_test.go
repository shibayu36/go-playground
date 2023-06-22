package main

import (
	"fmt"
	"testing"
	"time"
)

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()
	return ch
}

func TestGenerator(t *testing.T) {
	ch := generator("Hello!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func TestFanIn(t *testing.T) {
	ch1 := generator("Hello1!")
	ch2 := generator("Hello2!")
	ch := fanIn(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
