package main

import "sync"

type worker func()

func WorkerPool(jobs <-chan int, workers int, process func(int) int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(workers)

	workerFn := func() {
		defer wg.Done()
		for job := range jobs {
			out <- process(job)
		}
	}

	for range workers {
		go workerFn()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
