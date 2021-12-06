package clock

import (
	"sync/atomic"
	"time"

	"github.com/koykov/policy"
)

type sched struct {
	spinlock uint32
	lock     policy.Lock
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
	s.lock.Lock()
	defer s.lock.Unlock()
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
	defer s.sunlock()
	for i := 0; i < len(s.buf); i++ {
		r := &s.buf[i]
		if r.next.Before(now) {
			r.next = now.Add(r.dur)
			r.fn()
		}
	}
}
