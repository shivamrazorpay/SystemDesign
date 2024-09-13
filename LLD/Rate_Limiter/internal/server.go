package internal

import (
	"time"
)

func NewRateLimiter(windowSize time.Duration, maxRequests int) *RateLimiter {
	return &RateLimiter{
		WindowSize:    windowSize,
		MaxRequests:   maxRequests,
		RequestCounts: make(map[int64]int),
	}
}

func (rl *RateLimiter) AllowRequest() bool {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now().Unix()
	rl.cleanupOldRequests(now)

	// Calculate total requests in the current window
	totalRequests := 0
	for _, count := range rl.RequestCounts {
		totalRequests += count
	}

	if totalRequests >= rl.MaxRequests {
		// If we exceed the allowed number of requests, deny the request
		return false
	}

	// Allow the request and increment the request count for the current second
	rl.RequestCounts[now]++
	return true
}

func (rl *RateLimiter) cleanupOldRequests(currentTime int64) {
	// Remove any requests that fall outside the sliding window
	for timestamp := range rl.RequestCounts {
		if currentTime-timestamp >= int64(rl.WindowSize.Seconds()) {
			delete(rl.RequestCounts, timestamp)
		}
	}
}
