package main

import "sync"

type WaitGrp struct {
	count int
	cond  *sync.Cond
}

func NewWaitGrp() *WaitGrp {
	return &WaitGrp{
		count: 0,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *WaitGrp) Add(count int) {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	wg.count += count
}

func (wg *WaitGrp) Done() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	wg.count--
	if wg.count == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *WaitGrp) Wait() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	for wg.count > 0 {
		wg.cond.Wait()
	}
}
