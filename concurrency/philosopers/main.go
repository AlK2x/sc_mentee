package main

import (
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

type Table struct {
	fMutex [5]chan struct{}
}

func NewTable() *Table {
	t := &Table{}
	for i := range t.fMutex {
		t.fMutex[i] = make(chan struct{}, 1)
	}
	return t
}

func (t *Table) TakeLeftFork(seat int) {
	forkNum := seat
	t.fMutex[forkNum] <- struct{}{}
}

func (t *Table) TakeRightFork(seat int) {
	forkNum := (seat + 1) % 5
	t.fMutex[forkNum] <- struct{}{}
}

func (t *Table) ReturnLeftFork(seat int) {
	<-t.fMutex[seat]
}

func (t *Table) ReturnRightFork(seat int) {
	<-t.fMutex[(seat+1)%5]
}

func (t *Table) TakeFork(num int) bool {
	result := make(chan bool)
	go func() {
		select {
		case t.fMutex[num] <- struct{}{}:
			result <- true
		default:
			result <- false
		}
	}()
	return <-result
}

func (t *Table) ReturnFork(num int) {
	<-t.fMutex[num]
}

type PhilosopherState string

const (
	Thinking PhilosopherState = "thinking"
	Starving PhilosopherState = "starving"
	Eating   PhilosopherState = "eating"
)

type TakeForksStrategy interface {
	TakeForks(p *Philosopher)
	returnForks(p *Philosopher)
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
	eatCounter          int32
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
			p.starving()
			p.eating()
		}
	}
}

func (p *Philosopher) eating() {
	p.takeForks()
	p.state = Eating
	p.doAction()
	atomic.AddInt32(&p.eatCounter, 1)
	p.forksAccessStrategy.returnForks(p)
}

func (p *Philosopher) starving() {
	p.state = Starving
	p.doAction()
}

func (p *Philosopher) takeForks() {
	p.forksAccessStrategy.TakeForks(p)
}

func (p *Philosopher) thinking() {
	p.state = Thinking
	p.doAction()
}

func (p *Philosopher) doAction() {
	//p.printAction()
	sleep := rand.UintN(1000)
	time.Sleep(time.Millisecond * time.Duration(sleep))
}

func (p *Philosopher) CountEating() int {
	return int(atomic.LoadInt32(&p.eatCounter))
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
		if ps[i].CountEating() == 0 {
			fmt.Printf("Starving philosoper %d. Possible deadlock detected\n", ps[i].seat)
		}
	}
}
