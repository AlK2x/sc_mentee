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

func (ro *ResourceOrdering) returnForks(p *Philosopher) {
	ro.table.ReturnRightFork(p.seat)
	ro.table.ReturnLeftFork(p.seat)
}
