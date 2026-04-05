package main

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestFixedWindow(t *testing.T) {
	limit := 10
	window := 10 * time.Millisecond
	fw := NewFixedWindow(limit, window)
	for range 10 {
		if !fw.Allow() {
			t.Errorf("10 request should be successful")
		}
	}
	if fw.Allow() {
		t.Error("11-th request must be rejected")
	}
	time.Sleep(10 * time.Millisecond)
	for range 10 {
		if !fw.Allow() {
			t.Errorf("10 request should be successful after time window")
		}
	}
	if fw.Allow() {
		t.Error("11-th request must be rejected after time window")
	}
}

func TestRateLimiter(t *testing.T) {
	limit := 10
	window := 10 * time.Millisecond
	rl := NewRateLimiter(limit, window)
	clients := 3
	wg := sync.WaitGroup{}
	wg.Add(clients)
	for clientId := range clients {
		go func(clientId string) {
			defer wg.Done()

			for range 10 {
				if !rl.Allow(clientId) {
					t.Errorf("10 request should be successful")
				}
			}
			if rl.Allow(clientId) {
				t.Error("11-th request must be rejected")
			}
			time.Sleep(10 * time.Millisecond)
			for range 10 {
				if !rl.Allow(clientId) {
					t.Errorf("10 request should be successful after time window")
				}
			}
			if rl.Allow(clientId) {
				t.Error("11-th request must be rejected after time window")
			}

		}(strconv.Itoa(clientId))
	}
	wg.Wait()
}
