package main

import "sync"

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
