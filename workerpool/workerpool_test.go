package main

import (
	"reflect"
	"slices"
	"testing"
	"time"
)

func TestSignleWorker(t *testing.T) {
	jobs := make(chan int)
	input := []int{1, 2, 3}
	go func() {
		defer close(jobs)
		for _, val := range input {
			jobs <- val
		}
	}()
	expected := []int{1, 4, 9}

	out := WorkerPool(jobs, 1, func(i int) int { return i * i })

	assertReceiveFromChannel(t, out, expected, time.Second)
}

func TestWorkersMoreThanData(t *testing.T) {
	jobs := make(chan int)
	input := []int{1, 2}
	go func() {
		defer close(jobs)
		for _, val := range input {
			jobs <- val
		}
	}()
	expected := []int{1, 4}

	out := WorkerPool(jobs, 5, func(i int) int { return i * i })

	assertReceiveFromChannel(t, out, expected, time.Second)
}

func TestEmptyInput(t *testing.T) {
	jobs := make(chan int)
	input := []int{}
	go func() {
		defer close(jobs)
		for _, val := range input {
			jobs <- val
		}
	}()
	expected := []int{}

	out := WorkerPool(jobs, 5, func(i int) int { return i * i })

	assertReceiveFromChannel(t, out, expected, time.Second)
}

func TestSlowProcessFunction(t *testing.T) {
	jobs := make(chan int)
	input := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(jobs)
		for _, val := range input {
			jobs <- val
		}
	}()
	expected := []int{1, 4, 9, 16, 25}
	process := func(i int) int {
		time.Sleep(3 * time.Second)
		return i * i
	}

	out := WorkerPool(jobs, 5, process)

	assertReceiveFromChannel(t, out, expected, 4*time.Second)
}

func assertReceiveFromChannel(t *testing.T, ch <-chan int, expected []int, timeout time.Duration) {
	received := readDataFromChannel(t, ch, timeout)

	slices.Sort(received)
	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Received %v, got %v", received, expected)
	}
}

func readDataFromChannel(t *testing.T, ch <-chan int, timeout time.Duration) []int {
	received := make([]int, 0)
	timer := time.NewTimer(timeout)
	for {
		select {
		case <-timer.C:
			t.Fatalf("Workerpool timeout exceeded")
		case val, ok := <-ch:
			if !ok {
				return received
			}
			received = append(received, val)
		}
	}
}
