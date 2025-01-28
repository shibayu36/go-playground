package main

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"
)

type MyChannel[T any] struct {
	// 送信側のセマフォ。queueのsizeがmaxになったらブロックする
	sendSem *Semaphore
	// 受信側のセマフォ。queueが空の時にブロックする
	recvSem *Semaphore

	queueMu sync.Mutex
	queue   *list.List
}

func NewMyChannel[T any](size uint) *MyChannel[T] {
	return &MyChannel[T]{
		sendSem: NewSemaphore(size),
		recvSem: NewSemaphore(0),
		queue:   list.New(),
	}
}

func (c *MyChannel[T]) Send(v T) {
	c.sendSem.Acquire()
	c.queueMu.Lock()
	c.queue.PushBack(v)
	c.queueMu.Unlock()
	c.recvSem.Release()
}

func (c *MyChannel[T]) Recv() T {
	c.recvSem.Acquire()
	c.queueMu.Lock()
	val := c.queue.Remove(c.queue.Front())
	c.queueMu.Unlock()
	c.sendSem.Release()
	return val.(T)
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
