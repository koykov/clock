package clock

import (
	"sync"
	"sync/atomic"
	"time"
)

type sched struct {
	spinlock uint32
	mux      sync.RWMutex
	buf      []schedRule
}

type schedRule struct {
	fn   func()
	dur  time.Duration
	next time.Time
}

func (s *sched) slock() {
	atomic.StoreUint32(&s.spinlock, 1)
}

func (s *sched) sunlock() {
	atomic.StoreUint32(&s.spinlock, 0)
}

func (s *sched) slocked() bool {
	return atomic.LoadUint32(&s.spinlock) == 1
}

func (s *sched) register(dur time.Duration, fn func(), now time.Time) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.buf = append(s.buf, schedRule{
		fn:   fn,
		dur:  dur,
		next: now.Add(dur),
	})
}

func (s *sched) apply(now time.Time) {
	if len(s.buf) == 0 || s.slocked() {
		return
	}
	s.slock()
	s.mux.RLock()
	defer func() {
		s.mux.RUnlock()
		s.sunlock()
	}()
	for i := 0; i < len(s.buf); i++ {
		r := &s.buf[i]
		if r.next.Before(now) {
			r.next = now.Add(r.dur)
			r.fn()
		}
	}
}
