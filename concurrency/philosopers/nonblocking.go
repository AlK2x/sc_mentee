package main

import (
	"math/rand/v2"
	"time"
)

type NonBlockStrategy struct {
	table *Table
}

func (ro *NonBlockStrategy) TakeForks(p *Philosopher) {
	left := p.seat
	right := (p.seat + 1) % 5

	for {
		if ro.table.TakeFork(left) {
			if ro.table.TakeFork(right) {
				return
			} else {
				ro.table.ReturnFork(left)
			}
		}

		sleep := rand.UintN(1000)
		timer := time.NewTimer(time.Millisecond * time.Duration(sleep))
		<-timer.C
	}
}

func (ro *NonBlockStrategy) returnForks(p *Philosopher) {
	ro.table.ReturnLeftFork(p.seat)
	ro.table.ReturnRightFork(p.seat)
}
