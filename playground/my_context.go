package main

import (
	"sync"
	"sync/atomic"
	"time"
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

		// 子を全てcancelする
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
	// 親がcancelされることがないなら伝播を考える必要はない
	parentDone := c.parent.Done()
	if parentDone == nil {
		return
	}

	// すでに親がcancelされているならすぐさま自分もcancelする
	select {
	case <-parentDone:
		c.cancel(false)
		return
	default:
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
		defer goroutineCnt.Add(-1)
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

type myTimerCtx struct {
	myCancelCtx
	timer *time.Timer // myCancelCtxのmuを使って保護

	deadline time.Time
}

func NewMyContextWithDeadline(parent MyContext, deadline time.Time) (*myTimerCtx, func()) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	ctx := &myTimerCtx{
		myCancelCtx: myCancelCtx{
			parent:   parent,
			children: make(map[*myCancelCtx]struct{}),
			done:     make(chan struct{}),
		},
		deadline: deadline,
	}
	ctx.propagateCancel()

	duration := time.Until(deadline)
	if duration <= 0 {
		// すでに過ぎているならすぐにcancelする
		ctx.cancel(true)
		return ctx, func() { ctx.cancel(false) }
	}

	ctx.mu.Lock()
	ctx.timer = time.AfterFunc(duration, func() {
		// durationが過ぎたら自動でcancel
		ctx.cancel(true)
	})
	ctx.mu.Unlock()

	return ctx, func() { ctx.cancel(true) }
}

func (c *myTimerCtx) cancel(removeFromParent bool) {
	c.myCancelCtx.cancel(removeFromParent)

	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}
