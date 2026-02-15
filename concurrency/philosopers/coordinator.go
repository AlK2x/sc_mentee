package main

import "sync"

type Servant struct {
	m     sync.Mutex
	table *Table
}

func (s *Servant) TryTakeForks(seat int) bool {
	s.m.Lock()
	defer s.m.Unlock()
	s.table.TakeLeftFork(seat)
	s.table.TakeRightFork(seat)
	return false
}

func (s *Servant) ReturnForks(seat int) {
	s.table.ReturnLeftFork(seat)
	s.table.ReturnRightFork(seat)
}

type CentralCoordinator struct {
	servant *Servant
}

func (ro *CentralCoordinator) TakeForks(p *Philosopher) {
	for ro.servant.TryTakeForks(p.seat) {
	}
}

func (ro *CentralCoordinator) returnForks(p *Philosopher) {
	ro.servant.ReturnForks(p.seat)
}
