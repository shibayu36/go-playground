package main

import (
	"sync"
	"sync/atomic"
)

var goroutineCnt = atomic.Int64{}

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

type myCancelCtx struct {
	parent   MyContext
	children map[*myCancelCtx]struct{} // done伝播のためにchildrenを保持

	done     chan struct{}
	doneOnce sync.Once

	mu sync.Mutex
}

func NewMyContextWithCancel(parent MyContext) (*myCancelCtx, func()) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	done := make(chan struct{})
	ctx := &myCancelCtx{
		parent:   parent,
		children: make(map[*myCancelCtx]struct{}),
		done:     done,
	}
	ctx.propagateCancel()

	return ctx, func() { ctx.cancel(true) }
}

func (c *myCancelCtx) Done() <-chan struct{} {
	return c.done
}

func (c *myCancelCtx) cancel(removeFromParent bool) {
	c.doneOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		close(c.done)

		for child := range c.children {
			child.cancel(false)
		}
		c.children = nil

		if removeFromParent {
			// 親がMyContextWithCancelの場合は、親のchildrenから自分を削除して伝播の対象から外す
			parentCtx, ok := c.parent.(*myCancelCtx)
			if ok {
				parentCtx.removeChild(c)
			}
		}
	})
}

func (c *myCancelCtx) propagateCancel() {
	parentDone := c.parent.Done()
	if parentDone == nil {
		return
	}

	// 親がMyContextWithCancelの場合は、MyContextWithCancelのcancel側でchildrenに伝播するやり方に任せる
	if parentCtx, ok := c.parent.(*myCancelCtx); ok {
		parentCtx.mu.Lock()
		parentCtx.children[c] = struct{}{}
		parentCtx.mu.Unlock()
		return
	}

	// c.Done()での観察でしか伝播できない場合はgoroutineを起動して待つ
	goroutineCnt.Add(1)
	go func() {
		select {
		case <-parentDone:
			c.cancel(false)
		case <-c.Done():
			// 自分がキャンセルされたら親のDone待ちを抜ける
		}
	}()
}

func (c *myCancelCtx) removeChild(child *myCancelCtx) {
	c.mu.Lock()
	delete(c.children, child)
	c.mu.Unlock()
}
