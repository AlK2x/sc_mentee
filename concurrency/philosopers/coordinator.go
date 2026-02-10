package main

import "sync"

type Servant struct {
	m     sync.Mutex
	table *Table
}

func (s *Servant) TryTakeForks(seat int) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if s.table.TakeLeftFork(seat) {
		if s.table.TakeRightFork(seat) {
			return true
		}
	}
	s.table.ReturnForks(seat)
	return false
}

func (s *Servant) ReturnForks(seat int) {
	s.table.ReturnForks(seat)
}

type CentralCoordinator struct {
	servant *Servant
}

func (ro *CentralCoordinator) TakeForks(p *Philosopher) {
	for ro.servant.TryTakeForks(p.seat) {
	}
}

func (ro *CentralCoordinator) OnEatEnding(p *Philosopher) {
	ro.servant.ReturnForks(p.seat)
}
