package main

type ResourceOrdering struct {
	table *Table
}

func (ro *ResourceOrdering) TakeForks(p *Philosopher) {
	if p.seat == 4 {
		ro.table.TakeRightFork(p.seat)
		ro.table.TakeLeftFork(p.seat)
	} else {
		ro.table.TakeLeftFork(p.seat)
		ro.table.TakeRightFork(p.seat)
	}
}

func (ro *ResourceOrdering) OnEatEnding(p *Philosopher) {
	ro.table.ReturnForks(p.seat)
}
