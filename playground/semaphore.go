package main

import "sync"

type Semaphore struct {
	permits    uint       // 残っている許可数
	maxPermits uint       // 設定した許可数
	cond       *sync.Cond // 許可数が不足している時に待機する条件変数
}

func NewSemaphore(permits uint) *Semaphore {
	return &Semaphore{
		permits:    permits,
		maxPermits: permits,
		cond:       sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	for s.permits <= 0 {
		s.cond.Wait()
	}

	s.permits--
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	s.permits++
	// if s.permits < s.maxPermits {
	// 	s.permits++
	// }
	s.cond.Signal() // どれか一つに伝われば良い
}
