package main

type MyContext interface {
	Done() <-chan struct{}
}

type MyContextBackground struct{}

func (c *MyContextBackground) Done() <-chan struct{} {
	return nil
}

type MyContextWithCancel struct {
	done chan struct{}
}

func NewMyContextWithCancel() (*MyContextWithCancel, func()) {
	done := make(chan struct{})
	ctx := &MyContextWithCancel{
		done: done,
	}
	return ctx, ctx.cancel
}

func (c *MyContextWithCancel) Done() <-chan struct{} {
	return c.done
}

func (c *MyContextWithCancel) cancel() {
	close(c.done)
}
