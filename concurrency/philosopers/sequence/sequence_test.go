package main

import "testing"

func TestResourceOrdering(t *testing.T) {
	table := NewTable()
	ps := [5]*Philosopher{}
	for i := range len(ps) {
		p := Philosopher{
			seat:                i,
			table:               table,
			forksAccessStrategy: &ResourceOrdering{},
		}
		p.Start()
		ps[i] = &p
	}

	for i := range ps {
		if ps[i].CountEating() > 0 {
			t.Fatalf("Starving philosoper %d. Possible deadlock detected\n", ps[i].seat)
		}
	}
}
