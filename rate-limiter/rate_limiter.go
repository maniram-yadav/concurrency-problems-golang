package ratelimiter

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	capacity     int
	tokens       int
	fillInterval time.Duration
	tokenPerFill int
	lock         sync.Mutex
}

func NewRateLimiter(capacity, tokenPerFill int, interval time.Duration) *RateLimiter {

	if capacity <= 0 || tokenPerFill <= 0 || interval <= 0 {
		panic("Invalid parameters for rate limiter")
	}

	rateLimiter := &RateLimiter{
		capacity:     capacity,
		tokens:       capacity,
		tokenPerFill: tokenPerFill,
		fillInterval: interval,
	}
	go rateLimiter.startTokeRefill()
	return rateLimiter
}

func (r *RateLimiter) startTokeRefill() {
	ticker := time.NewTicker(r.fillInterval)
	for range ticker.C {
		fmt.Println("Inside loop for ticker fill")
		r.lock.Lock()
		if r.tokens < r.capacity {
			r.tokens = r.capacity
		}
		fmt.Println("Ticker have taken the lock")
		r.lock.Unlock()
	}
}

func (r *RateLimiter) Allow() bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	fmt.Println("Request have taken the lock")
	if r.tokens >= 0 {
		r.tokens--
		return true
	}
	return false
}
