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

type Clock struct {
	Precision time.Duration
	status    int32
	sec, nsec int64
	cancel    context.CancelFunc
}

var (
	_ = NewClock
)

func NewClock(prec time.Duration) *Clock {
	c := &Clock{Precision: prec}
	return c
}

func (c *Clock) Start() {
	if atomic.LoadInt32(&c.status) == StatusActive {
		return
	}
	atomic.StoreInt32(&c.status, StatusActive)
	if c.Precision == 0 {
		c.Precision = time.Second
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
		atomic.StoreInt32(&c.status, StatusIdle)
		c.cancel()
	}
}

func (c *Clock) Now() time.Time {
	return time.Unix(atomic.LoadInt64(&c.sec), atomic.LoadInt64(&c.nsec))
}

func (c *Clock) tick() {
	ts := time.Now().UnixNano()
	atomic.StoreInt64(&c.sec, ts/1e9)
	atomic.StoreInt64(&c.nsec, ts%1e9)
}
