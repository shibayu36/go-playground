package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

}

type myTestContext struct {
	MyContext

	done chan struct{}
}

func NewMyTestContext() *myTestContext {
	return &myTestContext{
		done: make(chan struct{}),
	}
}

func (c *myTestContext) Done() <-chan struct{} {
	return c.done
}

func (c *myTestContext) cancel() {
	close(c.done)
}

func Test_MyContext_PropagateCancel(t *testing.T) {
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

		assert.Equal(t, goroutineCnt.Load(), int64(0), "Done()待ちのgoroutineは起動していない")

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

	t.Run("独自Contextを待ってcancelされる", func(t *testing.T) {
		// - myTestContext
		//   - myCancelCtx
		//     - myCancelCtx
		//   - myCancelCtx

		rootTestCtx := NewMyTestContext()
		myCancelCtx1, _ := NewMyContextWithCancel(rootTestCtx)
		myCancelCtx1_2, _ := NewMyContextWithCancel(myCancelCtx1)
		myCancelCtx2, myCancelCtx2Cancel := NewMyContextWithCancel(rootTestCtx)

		assert.Equal(t, goroutineCnt.Load(), int64(2), "rootTestCtxを待つgoroutineのみ起動している")

		myCancelCtx2Cancel()

		select {
		case <-myCancelCtx2.Done():
			// goroutineが終了する確率を上げるために少し待つ
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, goroutineCnt.Load(), int64(1), "myCancelCtx2が待っていたgoroutineが終了した")
		case <-time.After(100 * time.Millisecond):
			t.Fatal("myCancelCtx2 was not cancelled")
		}

		rootTestCtx.cancel()

		// 全て伝播して終了する
		select {
		case <-myCancelCtx1.Done():
			// goroutineが終了する確率を上げるために少し待つ
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, goroutineCnt.Load(), int64(0), "全てのgoroutineが終了した")
		case <-time.After(100 * time.Millisecond):
			t.Fatal("myCancelCtx1 was not cancelled")
		}

		select {
		case <-myCancelCtx1_2.Done():
		case <-time.After(100 * time.Millisecond):
			t.Fatal("myCancelCtx1_2 was not cancelled")
		}
	})
}
