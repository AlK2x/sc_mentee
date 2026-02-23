package main

import (
	"sync"
)

func MergeChannels(chans ...<-chan int) <-chan int {
	result := make(chan int)

	var wg sync.WaitGroup
	readFn := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			result <- val
		}
	}

	wg.Add(len(chans))
	for _, ch := range chans {
		go readFn(ch)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
