package main

import (
	"fmt"
	"time"
)

const NUM_WORKERS = 5

func main() {
	jobs := make(chan int)
	go func() {
		defer close(jobs)
		for i := range 10 {
			jobs <- i
		}
	}()
	process := func(i int) int {
		time.Sleep(5 * time.Second) // emulate heavy work
		return i * i
	}
	out := WorkerPool(jobs, NUM_WORKERS, process)

	for val := range out {
		fmt.Printf("OUT: %d\n", val)
	}
}
