package backoff

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	for _, tt := range []struct {
		name    string
		base    time.Duration
		cap     time.Duration
		attempt int
		wantMin time.Duration
		wantMax time.Duration
	}{
		{
			name:    "ZeroBase",
			base:    0,
			cap:     time.Second,
			attempt: 0,
			wantMin: 0,
			wantMax: 0,
		},
		{
			name:    "ZeroCap",
			base:    time.Millisecond,
			cap:     0,
			attempt: 0,
			wantMin: 0,
			wantMax: 0,
		},
		{
			name:    "NegativeAttempt",
			base:    time.Millisecond,
			cap:     time.Second,
			attempt: -1,
			wantMin: 0,
			wantMax: 0,
		},
		{
			name:    "FirstAttempt",
			base:    100 * time.Millisecond,
			cap:     10 * time.Second,
			attempt: 0,
			wantMin: 0,
			wantMax: 100 * time.Millisecond,
		},
		{
			name:    "SecondAttempt",
			base:    100 * time.Millisecond,
			cap:     10 * time.Second,
			attempt: 1,
			wantMin: 0,
			wantMax: 200 * time.Millisecond,
		},
		{
			name:    "CappedByMaximum",
			base:    100 * time.Millisecond,
			cap:     300 * time.Millisecond,
			attempt: 3, // Would be 800ms without cap.
			wantMin: 0,
			wantMax: 300 * time.Millisecond,
		},
		{
			name:    "LargeAttemptNumber",
			base:    time.Millisecond,
			cap:     time.Second,
			attempt: 100, // Should be capped.
			wantMin: 0,
			wantMax: time.Second,
		},
		{
			name:    "LimitEqualsOne",
			base:    time.Nanosecond,
			cap:     time.Nanosecond,
			attempt: 0,
			wantMin: 0,
			wantMax: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// Test multiple times to check randomness behavior.
			for range 10 {
				got := Duration(tt.base, tt.cap, tt.attempt)
				if tt.wantMax == 0 {
					if got != 0 {
						t.Errorf("got %v, want 0", got)
					}
				} else {
					if got < tt.wantMin || got >= tt.wantMax {
						t.Errorf("got %v, want range [%v, %v)", got, tt.wantMin, tt.wantMax)
					}
				}
			}
		})
	}
}

func TestAfter(t *testing.T) {
	base := 10 * time.Millisecond
	cap := 50 * time.Millisecond
	attempt := 1
	wantMin := time.Duration(0)
	wantMax := 20 * time.Millisecond

	startTime := time.Now()
	ch := After(base, cap, attempt)
	if ch == nil {
		t.Fatal("unexpected nil")
	}

	select {
	case deliveredTime := <-ch:
		elapsed := time.Since(startTime)
		if deliveredTime.Before(startTime) {
			t.Error("got time before start, want time after start")
		}

		tolerance := 5 * time.Millisecond
		if elapsed < wantMin-tolerance {
			t.Errorf("got %v, want >= %v", elapsed, wantMin)
		}
		if elapsed > wantMax+tolerance {
			t.Errorf("got %v, want <= %v", elapsed, wantMax+tolerance)
		}
	case <-time.After(wantMax + 100*time.Millisecond):
		t.Error("got timeout, want timely delivery")
	}
}

func TestSleep(t *testing.T) {
	base := 5 * time.Millisecond
	cap := 20 * time.Millisecond
	attempt := 1
	wantMin := time.Duration(0)
	wantMax := 10 * time.Millisecond

	startTime := time.Now()
	Sleep(base, cap, attempt)
	elapsed := time.Since(startTime)

	tolerance := 5 * time.Millisecond
	if elapsed < wantMin {
		t.Errorf("got %v, want >= %v", elapsed, wantMin)
	}
	if elapsed > wantMax+tolerance {
		t.Errorf("got %v, want <= %v", elapsed, wantMax+tolerance)
	}
}
