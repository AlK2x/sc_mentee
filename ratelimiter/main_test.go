package main

import (
	"testing"
	"testing/synctest"
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

type testRequest struct {
	wait     time.Duration
	allowed  bool
	clientID string
}

func TestRateLimiter(t *testing.T) {
	cases := []struct {
		name     string
		limit    int
		window   time.Duration
		requests []testRequest
	}{
		{
			name:   "test flush time window",
			limit:  1,
			window: 1 * time.Second,
			requests: []testRequest{
				{0 * time.Second, true, "1"},
				{1 * time.Millisecond, false, "1"},
				{1 * time.Second, true, "1"},
				{1 * time.Second, true, "1"},
				{999 * time.Millisecond, false, "1"},
				{1 * time.Millisecond, true, "1"},
			},
		},
		{
			name:   "test multiple clients",
			limit:  1,
			window: 1 * time.Second,
			requests: []testRequest{
				{0 * time.Second, true, "1"},
				{0 * time.Second, true, "2"},
				{0 * time.Second, true, "3"},
				{0 * time.Second, false, "1"},
				{0 * time.Second, false, "2"},
				{0 * time.Second, false, "3"},
				{1 * time.Second, true, "1"}, // since 1 second time passed for other clients too
				{0 * time.Second, true, "2"},
				{0 * time.Second, true, "3"},
			},
		},
	}

	synctest.Test(t, func(t *testing.T) {
		for _, tt := range cases {
			rl := NewRateLimiter(tt.limit, tt.window)
			for i, r := range tt.requests {
				time.Sleep(r.wait)
				if rl.Allow(r.clientID) != r.allowed {
					t.Errorf("Test %s; Request %d Allow is wrong, expected %v", tt.name, i, r.allowed)
				}
			}
		}
	})
}

func TestRateLimiterConcurrent(t *testing.T) {
	clientID := "1"
	synctest.Test(t, func(t *testing.T) {
		rl := NewRateLimiter(5, 1*time.Second)
		allowedRequests := []testRequest{
			{0 * time.Second, true, clientID},
			{0 * time.Second, true, clientID},
			{0 * time.Second, true, clientID},
			{0 * time.Second, true, clientID},
			{0 * time.Second, true, clientID},
		}
		for i, r := range allowedRequests {
			go func() {
				if rl.Allow(r.clientID) != r.allowed {
					t.Errorf("Request %d Allow is wrong, expected %v", i, r.allowed)
				}
			}()
		}

		synctest.Wait()
		time.Sleep(999 * time.Millisecond)

		relectRequests := []testRequest{
			{0 * time.Second, false, clientID},
			{0 * time.Second, false, clientID},
			{0 * time.Second, false, clientID},
			{0 * time.Second, false, clientID},
			{0 * time.Second, false, clientID},
		}
		for i, r := range relectRequests {
			go func() {
				if rl.Allow(r.clientID) != r.allowed {
					t.Errorf("Request %d Allow is wrong, expected %v", i, r.allowed)
				}
			}()
		}
		synctest.Wait()
	})
}
