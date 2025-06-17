/*
Package backoff implements a Full-Jitter exponential backoff helper for Go.
*/
package backoff

import (
	"math/rand/v2"
	"time"
)

// Duration returns a randomized exponential-backoff delay. The delay is chosen
// uniformly from [0, min(cap, base*2^attempt)).
func Duration(base, cap time.Duration, attempt int) time.Duration {
	if base <= 0 || cap <= 0 || attempt < 0 {
		return 0
	}

	// Limit = base * 2^attempt, but never above cap and never overflow.
	var limit time.Duration
	if attempt >= 63 || base > cap>>attempt {
		limit = cap
	} else {
		limit = base << attempt
	}

	if limit <= 1 {
		return 0
	}
	return time.Duration(rand.N(int64(limit)))
}

// After returns a channel that will deliver the current time after the delay
// produced by [Duration]. It is shorthand for
// time.After(Duration(base, cap, attempt)).
func After(base, cap time.Duration, attempt int) <-chan time.Time {
	return time.After(Duration(base, cap, attempt))
}

// Sleep blocks for the delay produced by [Duration]. It is shorthand for
// time.Sleep(Duration(base, cap, attempt)).
func Sleep(base, cap time.Duration, attempt int) {
	time.Sleep(Duration(base, cap, attempt))
}
