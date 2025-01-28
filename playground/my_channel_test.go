package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type MyChannel[T any] struct {
	sendSem *Semaphore
	recvSem *Semaphore

	queueMu sync.Mutex
	queue   []T
}

func NewMyChannel[T any](size uint) *MyChannel[T] {
	return &MyChannel[T]{
		sendSem: NewSemaphore(size),
		recvSem: NewSemaphore(0),
		queue:   make([]T, 0, size),
	}
}

func (c *MyChannel[T]) Send(v T) {
	c.sendSem.Acquire()
	c.queueMu.Lock()
	c.queue = append(c.queue, v)
	c.queueMu.Unlock()
	c.recvSem.Release()
}

func (c *MyChannel[T]) Recv() T {
	c.recvSem.Acquire()
	c.queueMu.Lock()
	val := c.queue[0]
	c.queue = c.queue[1:]
	c.queueMu.Unlock()
	c.sendSem.Release()
	return val
}

func Test_MyChannel(t *testing.T) {
	c := NewMyChannel[int](1)
	go func() {
		c.Send(1)
		fmt.Println("Sended 1")
		time.Sleep(1 * time.Second)
		c.Send(2)
		fmt.Println("Sended 2")
		time.Sleep(1 * time.Second)
		c.Send(3)
		fmt.Println("Sended 3")
	}()
	fmt.Println("Recv: ", c.Recv())
	fmt.Println("Recv: ", c.Recv())
	fmt.Println("Recv: ", c.Recv())
}
