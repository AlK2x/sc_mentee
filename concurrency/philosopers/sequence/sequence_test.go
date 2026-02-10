package main

import (
	"testing"
	"time"
)

func TestPhilosoperProblem(t *testing.T) {
	cases := []struct {
		name string
		fn   func(table *Table) [5]*Philosopher
	}{
		{
			name: "Resource Ordering",
			fn: func(table *Table) [5]*Philosopher {
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
				return ps
			},
		},
		{
			name: "Restricted Parallelism",
			fn: func(table *Table) [5]*Philosopher {
				ps := [5]*Philosopher{}
				semaphore := NewSemaphore(4)
				for i := range len(ps) {
					p := Philosopher{
						seat:                i,
						table:               table,
						forksAccessStrategy: &RestrictParallelism{sem: semaphore},
					}
					p.Start()
					ps[i] = &p
				}
				return ps
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			table := NewTable()

			ps := tc.fn(table)

			timer := time.NewTimer(5 * time.Second)
			<-timer.C
			for i := range ps {
				if ps[i].CountEating() == 0 {
					t.Fatalf("Starving philosoper %d. Possible deadlock detected\n", ps[i].seat)
				}
			}
		})
	}
}
