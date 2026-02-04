package main

import (
	"fmt"
	"time"
)

type TokenBucketLimiter struct {
	tokens     chan struct{}
	ticker     *time.Ticker
	maxTokens  int
	refillRate time.Duration
}

func main() {
	// TODO: Implementare il rate limiter
	fmt.Println("Rate Limiter")
}

func NewTokenBucketLimiter(maxTokens int, refillRate time.Duration) *TokenBucketLimiter {
	tokens := make(chan struct{}, maxTokens)
	for i := 0; i < maxTokens; i++ {
		tokens <- struct{}{}
	}

	ticker := time.NewTicker(refillRate)
	rl := &TokenBucketLimiter{
		tokens:     tokens,
		ticker:     ticker,
		maxTokens:  maxTokens,
		refillRate: refillRate,
	}

	go func() {
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}()

	return rl
}

func (rl *TokenBucketLimiter) Wait() {
	<-rl.tokens
}

func (rl *TokenBucketLimiter) TryWait(timeout time.Duration) bool {
	select {
	case <-rl.tokens:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (rl *TokenBucketLimiter) Stop() {
	rl.ticker.Stop()
}
