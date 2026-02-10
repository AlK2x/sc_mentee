package main

import (
	"testing"
	"time"
)

func TestPhilosoperProblem(t *testing.T) {
	cases := []struct {
		name           string
		runPhilosopers func(table *Table) [5]*Philosopher
	}{
		{
			name: "Resource Ordering",
			runPhilosopers: func(table *Table) [5]*Philosopher {
				ps := [5]*Philosopher{}
				for i := range len(ps) {
					p := Philosopher{
						seat:                i,
						forksAccessStrategy: &ResourceOrdering{table: table},
					}
					p.Start()
					ps[i] = &p
				}
				return ps
			},
		},
		{
			name: "Restricted Parallelism",
			runPhilosopers: func(table *Table) [5]*Philosopher {
				ps := [5]*Philosopher{}
				semaphore := NewSemaphore(4)
				for i := range len(ps) {
					p := Philosopher{
						seat:                i,
						forksAccessStrategy: &RestrictParallelism{table: table, sem: semaphore},
					}
					p.Start()
					ps[i] = &p
				}
				return ps
			},
		},
		{
			name: "Central Coordinator",
			runPhilosopers: func(table *Table) [5]*Philosopher {
				ps := [5]*Philosopher{}
				servant := &Servant{table: table}
				for i := range len(ps) {
					p := Philosopher{
						seat:                i,
						forksAccessStrategy: &CentralCoordinator{servant: servant},
					}
					p.Start()
					ps[i] = &p
				}
				return ps
			},
		},
		{
			name: "Non Blocking",
			runPhilosopers: func(table *Table) [5]*Philosopher {
				ps := [5]*Philosopher{}
				for i := range len(ps) {
					p := Philosopher{
						seat:                i,
						forksAccessStrategy: &NonBlockStrategy{table: table},
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

			ps := tc.runPhilosopers(table)

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
