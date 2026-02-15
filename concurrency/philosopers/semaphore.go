package main

import "sync"

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		size: n,
	}
}

type Semaphore struct {
	n    int
	size int
	mu   sync.Mutex
}

func (s *Semaphore) TryAcquare() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.n+1 < s.size {
		s.n += 1
		return true
	}
	return false
}

func (s *Semaphore) Release() {
	s.mu.Lock()
	s.n -= 1
	s.mu.Unlock()
}

func NewRestrictedParallelism(sem *Semaphore) RestrictParallelism {
	return RestrictParallelism{
		sem: sem,
	}
}

type RestrictParallelism struct {
	table *Table
	sem   *Semaphore
}

func (ro *RestrictParallelism) TakeForks(p *Philosopher) {
	for !ro.sem.TryAcquare() {
	}

	ro.table.TakeLeftFork(p.seat)
	ro.table.TakeRightFork(p.seat)
}

func (ro *RestrictParallelism) returnForks(p *Philosopher) {
	ro.table.ReturnLeftFork(p.seat)
	ro.table.ReturnRightFork(p.seat)
	ro.sem.Release()
}
