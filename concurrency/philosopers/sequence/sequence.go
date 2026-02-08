package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

type Table struct {
	forks  [5]bool
	fMutex [5]sync.Mutex
}

func NewTable() *Table {
	return &Table{
		forks: [5]bool{true, true, true, true, true},
	}
}

func (t *Table) TakeLeftFork(seat int) bool {
	if seat < 0 || seat >= 5 {
		return false
	}
	t.fMutex[seat].Lock()
	defer t.fMutex[seat].Unlock()
	if t.forks[seat] {
		t.forks[seat] = false
		return true
	}
	return false
}

func (t *Table) TakeRightFork(seat int) bool {
	if seat < 0 || seat >= 5 {
		return false
	}

	s := (seat + 1) % 5
	t.fMutex[s].Lock()
	defer t.fMutex[s].Unlock()
	if t.forks[s] {
		t.forks[s] = false
		return true
	}
	return false
}

func (t *Table) ReturnForks(seat int) {
	s := (seat + 1) % 5
	t.fMutex[seat].Lock()
	t.fMutex[s].Lock()

	t.forks[seat] = true
	t.forks[s] = true

	t.fMutex[s].Unlock()
	t.fMutex[seat].Unlock()
}

type PhilosopherState string

const (
	Thinking PhilosopherState = "thinking"
	Starving PhilosopherState = "starving"
	Eating   PhilosopherState = "eating"
)

type TakeForksStrategy interface {
	TakeForks(p *Philosopher)
}

type ResourceOrdering struct {
}

func (ro *ResourceOrdering) TakeForks(p *Philosopher) {
	if p.seat == 4 {
		p.table.TakeRightFork(p.seat)
		p.table.TakeLeftFork(p.seat)
	} else {
		p.table.TakeLeftFork(p.seat)
		p.table.TakeRightFork(p.seat)
	}
	p.state = Eating
	atomic.AddUint32(&p.eatCounter, 1)
}

type Philosopher struct {
	state               PhilosopherState
	seat                int
	table               *Table
	prevState           PhilosopherState
	forksAccessStrategy TakeForksStrategy
	eatCounter          uint32
}

func (p *Philosopher) Start() {
	p.state = Thinking
	go p.doStart()
}

func (p *Philosopher) doStart() {
	for {
		p.thinking()

		p.takeForks()

		p.eating()
		p.table.ReturnForks(p.seat)
	}
}

func (p *Philosopher) takeForks() {
	p.forksAccessStrategy.TakeForks(p)
}

func (p *Philosopher) eating() {
	p.state = Eating
	p.doAction()
}

func (p *Philosopher) thinking() {
	p.state = Thinking
	p.doAction()
}

func (p *Philosopher) doAction() {
	//p.printAction()
	sleep := rand.UintN(10)
	time.Sleep(time.Microsecond * time.Duration(sleep))
}

func (p *Philosopher) CountEating() int {
	return int(atomic.LoadUint32(&p.eatCounter))
}

func (p *Philosopher) printAction() {
	if p.state == p.prevState {
		return
	}
	p.prevState = p.state
	fmt.Printf("%d-th philosoper is %s\n", p.seat, p.state)
}

func main() {
	table := NewTable()
	n := 5

	ps := [5]*Philosopher{}
	for i := range n {
		p := Philosopher{
			seat:                i,
			table:               table,
			forksAccessStrategy: &ResourceOrdering{},
		}
		p.Start()
		ps[i] = &p
	}

	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	for i := range ps {
		if ps[i].CountEating() > 0 {
			fmt.Printf("Starving philosoper %d. Possible deadlock detected\n", ps[i].seat)
		}
	}
}
