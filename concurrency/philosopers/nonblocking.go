package main

type NonBlockStrategy struct {
	table *Table
}

func (ro *NonBlockStrategy) TakeForks(p *Philosopher) {
	ro.table.TakeLeftFork(p.seat)
	ro.table.TakeRightFork(p.seat)
}

func (ro *NonBlockStrategy) OnEatEnding(p *Philosopher) {
	ro.table.ReturnForks(p.seat)
}
