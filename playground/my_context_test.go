package main

import (
	"testing"
	"time"
)

func Test_MyContext(t *testing.T) {
	ctx, cancel := NewMyContextWithCancel()
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()

	select {
	case <-ctx.Done():
		// 一定時間後にctx.Done()を抜けられる
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout")
	}
}
