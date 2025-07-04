package mdwratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	Interval time.Duration
	Rw       sync.RWMutex
	Request  sync.Map
	Limit    int
}

// limit request, long interval for request
func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		Interval: interval,
		Limit:    limit,
	}
}
