package main

import (
	"sync"
	"testing"
	"time"
)

type ReadWriteLock struct {
	readersLock sync.Mutex
	readerCount uint
	globalLock  sync.Mutex
}

func (l *ReadWriteLock) ReadLock() {
	l.readersLock.Lock()
	l.readerCount++
	if l.readerCount == 1 {
		l.globalLock.Lock()
	}
	l.readersLock.Unlock()
}

func (l *ReadWriteLock) ReadUnlock() {
	l.readersLock.Lock()
	l.readerCount--
	if l.readerCount == 0 {
		l.globalLock.Unlock()
	}
	l.readersLock.Unlock()
}

func (l *ReadWriteLock) WriteLock() {
	l.globalLock.Lock()
}

func (l *ReadWriteLock) WriteUnlock() {
	l.globalLock.Unlock()
}

func TestReadWriteLock(t *testing.T) {
	t.Run("複数のReadLockを呼び出した後、WriteLockはブロックされる。ReadLockが全て終わったらWriteLockが呼び出される", func(t *testing.T) {
		l := &ReadWriteLock{}
		readStarted := make(chan struct{})
		readersReady := make(chan struct{})
		allowReadersToFinish := make(chan struct{})

		// 複数のゴルーチンがReadLockを取得できる
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				l.ReadLock()
				readStarted <- struct{}{} // 読み込みロックを取得したことを通知
				<-allowReadersToFinish    // 続行の許可を待つ
				l.ReadUnlock()
			}()
		}

		// すべての読み込みゴルーチンが開始されるのを待つ
		for i := 0; i < 10; i++ {
			<-readStarted
		}
		close(readersReady)

		// WriteLockがブロックされることを確認
		writeBlocked := make(chan struct{})
		go func() {
			l.WriteLock()
			close(writeBlocked)
			l.WriteUnlock()
		}()

		// WriteLockがブロックされていることを確認
		select {
		case <-writeBlocked:
			t.Error("WriteLock should be blocked")
		default:
		}

		// 読み込みゴルーチンの終了を許可
		close(allowReadersToFinish)
		wg.Wait()

		// WriteLockが取得できることを確認
		select {
		case <-writeBlocked:
			// 期待通り取得できた
		case <-time.After(100 * time.Millisecond):
			t.Error("WriteLock should be acquired")
		}
	})
}
