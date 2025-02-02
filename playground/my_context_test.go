package main

import (
	"testing"
	"time"
)

func Test_MyContext(t *testing.T) {

	t.Run("500ms後にcancel", func(t *testing.T) {
		ctx, cancel := NewMyContextWithCancel(NewMyContextBackground())
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
	})

	t.Run("親がcancelされたら子もcancelされる", func(t *testing.T) {
		parent, cancel := NewMyContextWithCancel(NewMyContextBackground())
		child1, _ := NewMyContextWithCancel(parent)
		child2, _ := NewMyContextWithCancel(parent)

		go func() {
			time.Sleep(100 * time.Millisecond)
			cancel()
		}()

		// 全てのコンテキストのDone()チャネルがcloseされることを確認
		select {
		case <-parent.Done():
			// parentがキャンセルされたことを確認
		case <-time.After(1 * time.Second):
			t.Fatal("parent context was not cancelled")
		}

		select {
		case <-child1.Done():
			// child1がキャンセルされたことを確認
		case <-time.After(100 * time.Millisecond):
			t.Fatal("child1 context was not cancelled")
		}

		select {
		case <-child2.Done():
			// child2がキャンセルされたことを確認
		case <-time.After(100 * time.Millisecond):
			t.Fatal("child2 context was not cancelled")
		}
	})
}
