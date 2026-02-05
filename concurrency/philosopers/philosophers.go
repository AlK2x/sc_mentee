package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Table struct {
	forks [5]bool
}

func NewTable() *Table {
	return &Table{
		forks: [5]bool{true, true, true, true, true},
	}
}

func (t *Table) TakeLeftFork(seat int) bool {
	if t.forks[seat] {
		t.forks[seat] = false
		return true
	}
	return false
}

func (t *Table) TakeRightFork(seat int) bool {
	s := (seat + 1) % 5
	if t.forks[s] {
		t.forks[s] = false
		return true
	}
	return false
}

func (t *Table) PlaceFork(seat int) {
	t.forks[seat] = true
	t.forks[(seat+1)%5] = true
}

type PhilosopherState string

const (
	Thinking PhilosopherState = "thinking"
	Starving PhilosopherState = "starving"
	Eating   PhilosopherState = "eating"
)

type Philosopher struct {
	state     PhilosopherState
	seat      int
	table     *Table
	prevState PhilosopherState
}

func (p *Philosopher) Start() {
	p.state = Thinking
	go p.doStart()
}

func (p *Philosopher) doStart() {
	for {
		p.doAction()
		switch p.state {
		case Starving:
			if p.table.TakeLeftFork(p.seat) && p.table.TakeRightFork(p.seat) {
				p.state = Eating
			}
		case Thinking:
			p.state = Starving
		case Eating:
			p.state = Thinking
			p.table.PlaceFork(p.seat)
		}
	}
}

func (p *Philosopher) doAction() {
	p.printNewAction()
	sleep := rand.UintN(5)
	time.Sleep(time.Second * time.Duration(sleep))
}

func (p *Philosopher) printNewAction() {
	if p.state == p.prevState {
		return
	}
	p.prevState = p.state
	fmt.Printf("%d-th philosoper is %s\n", p.seat, p.state)
}

type Servant struct {
	m     sync.Mutex
	table *Table
}

func (s *Servant) GiveLeftFork(seat int) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if s.table.TakeLeftFork(seat) {
		if s.table.TakeRightFork(seat) {
			return true
		} else {
			s.table.PlaceFork(seat)
		}
	}
	return false
}

func (s *Servant) GiveRightFork(seat int) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if s.table.TakeLeftFork(seat) {
		if s.table.TakeRightFork(seat) {
			return true
		} else {
			s.table.PlaceFork(seat)
		}
	}
	return false
}

func main() {
	table := NewTable()
	n := 5
	for i := range n {
		p := Philosopher{seat: i, table: table}
		p.Start()
	}
	ch := make(chan struct{})
	<-ch
}
