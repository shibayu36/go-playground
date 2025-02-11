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

	t.Run("すでに親がcancelされているならすぐに自分もcancelされる", func(t *testing.T) {
		parent, cancel := NewMyContextWithCancel(NewMyContextBackground())
		cancel()

		child, _ := NewMyContextWithCancel(parent)

		select {
		case <-child.Done():
			// childがキャンセルされたことを確認
		default:
			t.Fatal("child context was not cancelled")
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

func Test_MyContext_WithDeadline(t *testing.T) {
	t.Run("deadlineが来たらcancelされる", func(t *testing.T) {
		ctx, _ := NewMyContextWithDeadline(NewMyContextBackground(), time.Now().Add(100*time.Millisecond))

		select {
		case <-ctx.Done():
			t.Fatal("ctx.Done() should be blocked")
		default:
		}

		time.Sleep(110 * time.Millisecond)

		select {
		case <-ctx.Done():
			// キャンセルされた
		default:
			t.Fatal("ctx.Done() should be closed")
		}
	})

	t.Run("すでにdeadlineが来ているならすぐにcancelされる", func(t *testing.T) {
		ctx, _ := NewMyContextWithDeadline(NewMyContextBackground(), time.Now().Add(-100*time.Millisecond))

		select {
		case <-ctx.Done():
			// キャンセルされた
		default:
			t.Fatal("ctx.Done() should be closed")
		}
	})

	t.Run("cancelを手動で呼ぶとcancelされる", func(t *testing.T) {
		ctx, cancel := NewMyContextWithDeadline(NewMyContextBackground(), time.Now().Add(100*time.Millisecond))

		select {
		case <-ctx.Done():
			t.Fatal("ctx.Done() should be blocked")
		default:
		}

		cancel()

		select {
		case <-ctx.Done():
			// キャンセルされた
		default:
			t.Fatal("ctx.Done() should be closed")
		}

		// その後、Deadlineが来てもエラーにならない
		time.Sleep(110 * time.Millisecond)
		select {
		case <-ctx.Done():
			// キャンセルされた
		default:
			t.Fatal("ctx.Done() should be closed")
		}
	})

	t.Run("親がキャンセルされたら子もキャンセルされる", func(t *testing.T) {
		parent, cancel := NewMyContextWithDeadline(NewMyContextBackground(), time.Now().Add(50*time.Millisecond))
		child, _ := NewMyContextWithDeadline(parent, time.Now().Add(1000*time.Millisecond))

		select {
		case <-parent.Done():
			t.Fatal("parent context was not blocked")
		case <-child.Done():
			t.Fatal("child context was not blocked")
		default:
		}

		cancel()

		select {
		case <-child.Done():
			// childがキャンセルされた
		default:
			t.Fatal("child context was not cancelled")
		}
	})

	t.Run("myTimerCtxを待つ時、goroutineが起動しない", func(t *testing.T) {
		goroutineCnt.Store(0)

		myTimerCtx, myTimerCtxCancel := NewMyContextWithDeadline(NewMyContextBackground(), time.Now().Add(100*time.Millisecond))
		myCancelCtx1, _ := NewMyContextWithCancel(myTimerCtx)
		myCancelCtx2, _ := NewMyContextWithCancel(myTimerCtx)

		select {
		case <-myTimerCtx.Done():
			t.Fatal("myTimerCtx was not blocked")
		case <-myCancelCtx1.Done():
			t.Fatal("myCancelCtx1 was not blocked")
		case <-myCancelCtx2.Done():
			t.Fatal("myCancelCtx2 was not blocked")
		default:
		}

		assert.Equal(t, int64(0), goroutineCnt.Load(), "goroutineが起動していない")

		myTimerCtxCancel()

		select {
		case <-myTimerCtx.Done():
			// キャンセルされた
		default:
			t.Fatal("myTimerCtx was not cancelled")
		}

		select {
		case <-myCancelCtx1.Done():
			// 子もキャンセルされた
		default:
			t.Fatal("myCancelCtx1 was not cancelled")
		}

		select {
		case <-myCancelCtx2.Done():
			// 子もキャンセルされた
		default:
			t.Fatal("myCancelCtx2 was not cancelled")
		}
	})
}
