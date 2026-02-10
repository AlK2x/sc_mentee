package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
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
	forkNum := seat
	t.fMutex[forkNum].Lock()
	defer t.fMutex[forkNum].Unlock()
	if t.forks[forkNum] {
		t.forks[forkNum] = false
		return true
	}
	return false
}

func (t *Table) TakeRightFork(seat int) bool {
	if seat < 0 || seat >= 5 {
		return false
	}

	forkNum := (seat + 1) % 5
	t.fMutex[forkNum].Lock()
	defer t.fMutex[forkNum].Unlock()
	if t.forks[forkNum] {
		t.forks[forkNum] = false
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
	OnEatEnding(p *Philosopher)
}

func NewPhilosoper(seat int, strategy TakeForksStrategy) *Philosopher {
	return &Philosopher{
		state:               Thinking,
		seat:                seat,
		forksAccessStrategy: strategy,
		stop:                make(chan struct{}),
	}
}

type Philosopher struct {
	state               PhilosopherState
	seat                int
	prevState           PhilosopherState
	forksAccessStrategy TakeForksStrategy
	eatCounter          int
	stop                chan struct{}
}

func (p *Philosopher) Start() {
	go p.doStart()
}

func (p *Philosopher) Stop() {
	p.stop <- struct{}{}
}

func (p *Philosopher) doStart() {
	for {
		select {
		case <-p.stop:
			return
		default:
			p.thinking()
			p.eating()
		}
	}
}

func (p *Philosopher) eating() {
	p.takeForks()
	p.state = Eating
	p.doAction()
	p.eatCounter++
	p.forksAccessStrategy.OnEatEnding(p)
}

func (p *Philosopher) takeForks() {
	p.forksAccessStrategy.TakeForks(p)
}

func (p *Philosopher) thinking() {
	p.state = Thinking
	p.doAction()
	p.state = Starving
}

func (p *Philosopher) doAction() {
	//p.printAction()
	sleep := rand.UintN(10)
	time.Sleep(time.Millisecond * time.Duration(sleep))
}

func (p *Philosopher) CountEating() int {
	return p.eatCounter
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
	strategy := &ResourceOrdering{table}
	for i := range n {
		p := NewPhilosoper(i, strategy)
		p.Start()
		ps[i] = p
	}

	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	for i := range ps {
		ps[i].Stop()
		if ps[i].CountEating() == 0 {
			fmt.Printf("Starving philosoper %d. Possible deadlock detected\n", ps[i].seat)
		}
	}
}
