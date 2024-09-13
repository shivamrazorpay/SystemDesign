package internal

import (
	"sync"
	"time"
)

type RateLimiter struct {
	sync.Mutex
	WindowSize    time.Duration // Duration of the sliding window
	MaxRequests   int           // Maximum number of requests allowed within the window
	RequestCounts map[int64]int // Keeps track of requests count in each second of the window
}
