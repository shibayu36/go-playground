package main

import (
	"testing"
)

func Test_MyContext(t *testing.T) {
	t.Run("500ms後にcancel", func(t *testing.T) {
		ctx, cancel := NewMyContextWithCancel(NewMyContextBackground())

		select {
		case <-ctx.Done():
			t.Fatal("ctx.Done() should be blocked")
		default:
		}

		cancel()

		select {
		case <-ctx.Done():
			// 一定時間後にctx.Done()を抜けられる
		default:
			t.Fatal("Timeout")
		}
	})

	t.Run("親がcancelされたら子もcancelされる", func(t *testing.T) {
		parent, cancel := NewMyContextWithCancel(NewMyContextBackground())
		child1, _ := NewMyContextWithCancel(parent)
		child2, _ := NewMyContextWithCancel(parent)

		select {
		case <-parent.Done():
			t.Fatal("parent context was not blocked")
		case <-child1.Done():
			t.Fatal("child1 context was not blocked")
		case <-child2.Done():
			t.Fatal("child2 context was not blocked")
		default:
		}

		cancel()

		// 全てのコンテキストのDone()チャネルがcloseされることを確認
		select {
		case <-parent.Done():
			// parentがキャンセルされた
		default:
			t.Fatal("parent context was not cancelled")
		}

		select {
		case <-child1.Done():
			// child1がキャンセルされたことを確認
		default:
			t.Fatal("child1 context was not cancelled")
		}

		select {
		case <-child2.Done():
			// child2がキャンセルされたことを確認
		default:
			t.Fatal("child2 context was not cancelled")
		}
	})
}
