package test

import (
	ratelimiter "concurrency-patterns/rate-limiter"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {

	rateLimiter := ratelimiter.NewRateLimiter(3, 1, time.Second)
	t.Log("rate limiter setup done")
	for i := 1; i <= 5; i++ {
		if rateLimiter.Allow() {
			t.Logf("Test: Request %d allowed\n", i)
		} else {
			t.Logf("Test: Request %d denied\n", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
