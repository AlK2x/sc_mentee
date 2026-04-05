package main

import "sync"

type worker func()

func WorkerPool(jobs <-chan int, workers int, process func(int) int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	workerFn := func() {
		defer wg.Done()
		for job := range jobs {
			out <- process(job)
		}
	}

	wg.Add(workers)
	for range workers {
		go workerFn()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
