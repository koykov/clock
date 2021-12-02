package clock

import (
	"context"
	"sync/atomic"
	"time"
)

const (
	StatusIdle int32 = iota
	StatusActive
)

// Clock is a fast replacement of base methods of time package.
type Clock struct {
	// Clock precision.
	// Settings this param too small (less than microseconds) or too big (great than second) is counterproductive.
	Precision time.Duration

	status int32
	sec, nsec,
	delta int64

	sched *sched

	cancel context.CancelFunc
}

// NewClock makes new clock with default precision (1 millisecond).
func NewClock() *Clock {
	return NewClockWP(time.Millisecond)
}

// NewClockWP makes new clock with given precision.
func NewClockWP(precision time.Duration) *Clock {
	c := &Clock{Precision: precision}
	return c
}

// Start initializes and starts the clock.
func (c *Clock) Start() {
	if atomic.LoadInt32(&c.status) == StatusActive {
		return
	}
	atomic.StoreInt32(&c.status, StatusActive)
	if c.Precision == 0 {
		c.Precision = time.Millisecond
	}
	c.tick()
	var ctx context.Context
	ctx, c.cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		t := time.NewTicker(c.Precision)
		for {
			select {
			case <-t.C:
				c.tick()
			case <-ctx.Done():
				t.Stop()
				return
			}
		}
	}(ctx)
}

func (c *Clock) Stop() {
	if atomic.LoadInt32(&c.status) == StatusActive {
		c.cancel()
		atomic.StoreInt32(&c.status, StatusIdle)
	}
}

// Now returns current time.
func (c *Clock) Now() time.Time {
	return time.Unix(atomic.LoadInt64(&c.sec), atomic.LoadInt64(&c.nsec))
}

// Jump performs time travel.
func (c *Clock) Jump(delta time.Duration) {
	atomic.AddInt64(&c.delta, int64(delta))
	c.tick()
}

func (c *Clock) Relative(raw string) time.Time {
	if dur, err := Relative(raw); err == nil {
		return c.Now().Add(dur)
	}
	return time.Time{}
}

func (c *Clock) Schedule(dur time.Duration, fn func()) {
	if c.sched == nil {
		c.sched = &sched{}
	}
	c.sched.register(dur, fn, c.Now())
}

func (c *Clock) tick() {
	ts := time.Now().UnixNano() + atomic.LoadInt64(&c.delta)
	atomic.StoreInt64(&c.sec, ts/1e9)
	atomic.StoreInt64(&c.nsec, ts%1e9)
	if c.sched != nil {
		c.sched.apply(c.Now())
	}
}
