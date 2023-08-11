package main

import (
	"context"
	"testing"
	"time"
)

func TestNestedContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a child context
	childCtx, childCancel := context.WithCancel(ctx)
	defer childCancel()

	go func() {
		time.Sleep(time.Millisecond * 100 * 2)
		cancel()
	}()

	go func() {
		time.Sleep(time.Millisecond * 100 * 3)
		// time.Sleep(time.Millisecond * 100 * 1)
		childCancel()
	}()

	parentCanceled, childCanceled := false, false

	for {
		if parentCanceled && childCanceled {
			t.Log("Both contexts are canceled")
			break
		}

		select {
		case <-childCtx.Done():
			if !childCanceled {
				t.Log("Child context is done")
				childCanceled = true
			}
		case <-ctx.Done():
			if !parentCanceled {
				t.Log("Parent context is done")
				parentCanceled = true
			}
		}
	}
}
