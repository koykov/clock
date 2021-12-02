package clock

import (
	"sync/atomic"
	"time"
)

type sched struct {
	spinlock uint32
	buf      []schedRule
}

type schedRule struct {
	fn   func()
	dur  time.Duration
	next time.Time
}

func (s *sched) lock() {
	atomic.StoreUint32(&s.spinlock, 1)
}

func (s *sched) unlock() {
	atomic.StoreUint32(&s.spinlock, 0)
}

func (s *sched) locked() bool {
	return atomic.LoadUint32(&s.spinlock) == 1
}

func (s *sched) register(dur time.Duration, fn func(), now time.Time) {
	s.buf = append(s.buf, schedRule{
		fn:   fn,
		dur:  dur,
		next: now.Add(dur),
	})
}

func (s *sched) apply(now time.Time) {
	if s.locked() || len(s.buf) == 0 {
		return
	}
	s.lock()
	defer s.unlock()
	for i := 0; i < len(s.buf); i++ {
		if s.buf[i].next.Before(now) {
			r := &s.buf[i]
			r.next = now.Add(r.dur)
			r.fn()
		}
	}
}
