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
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return ch
}

func generatorWithQuit(msg string, quit <-chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Millisecond * 100)
			case <-quit:
				fmt.Println("Quit!")
				return
			}
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

func fanInBySelect(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				ch <- s
			case s := <-ch2:
				ch <- s
			}
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

func TestFanInBySelect(t *testing.T) {
	ch1 := generator("Hello1!")
	ch2 := generator("Hello2!")
	ch := fanInBySelect(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func TestTimeoutLoop(t *testing.T) {
	ch := generator("Hello!")
	for i := 0; i < 5; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1200 * time.Millisecond):
			// case <-time.After(800 * time.Millisecond):
			return
		}
	}
}

func TestOverallTimeout(t *testing.T) {
	ch := generator("Hello!")
	timeout := time.After(3 * time.Millisecond * 100)
	for i := 0; i < 5; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			return
		}
	}
}

func TestOverallTimeoutWithQuit(t *testing.T) {
	quit := make(chan bool)
	ch := generatorWithQuit("Hello!", quit)
	timeout := time.After(3 * time.Millisecond * 100)
	for i := 0; i < 5; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			quit <- true
			return
		}
	}
}
