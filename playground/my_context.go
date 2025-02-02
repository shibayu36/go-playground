package main

import (
	"sync"
)

type MyContext interface {
	Done() <-chan struct{}
}

type MyContextBackground struct{}

func NewMyContextBackground() *MyContextBackground {
	return &MyContextBackground{}
}

func (c *MyContextBackground) Done() <-chan struct{} {
	return nil
}

type MyContextWithCancel struct {
	done     chan struct{}
	doneOnce sync.Once
}

func NewMyContextWithCancel(parent MyContext) (*MyContextWithCancel, func()) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	done := make(chan struct{})
	ctx := &MyContextWithCancel{
		done: done,
	}
	ctx.propagateCancel(parent)

	return ctx, ctx.cancel
}

func (c *MyContextWithCancel) Done() <-chan struct{} {
	return c.done
}

func (c *MyContextWithCancel) cancel() {
	c.doneOnce.Do(func() {
		close(c.done)
	})
}

func (c *MyContextWithCancel) propagateCancel(parent MyContext) {
	parentDone := parent.Done()
	if parentDone == nil {
		return
	}

	go func() {
		select {
		case <-parentDone:
			c.cancel()
		case <-c.Done():
			// 自分がキャンセルされたら親のDone待ちを抜ける
		}
	}()
}
