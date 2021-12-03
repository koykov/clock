package clock

import (
	"sync"
	"sync/atomic"
	"time"
)

type sched struct {
	spinlock uint32
	idx      uint32
	buf      [2][]schedRule

	// Pseudo lock staff.
	flag uint32
	mux  sync.Mutex
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
	s.lock()
	defer s.unlock()
	act, opp := s.act(), s.opp()
	s.buf[opp] = append(s.buf[opp][:0], s.buf[act]...)
	s.buf[opp] = append(s.buf[opp], schedRule{
		fn:   fn,
		dur:  dur,
		next: now.Add(dur),
	})
	s.rot()
}

func (s *sched) apply(now time.Time) {
	buf := s.getBuf()
	if len(buf) == 0 || s.slocked() {
		return
	}
	s.slock()
	defer s.sunlock()
	for i := 0; i < len(buf); i++ {
		r := &buf[i]
		if r.next.Before(now) {
			r.next = now.Add(r.dur)
			r.fn()
		}
	}
}

func (s *sched) act() uint32 {
	return atomic.LoadUint32(&s.idx)
}

func (s *sched) opp() uint32 {
	act := s.act()
	if act == 0 {
		return 1
	}
	return 0
}

func (s *sched) rot() {
	if act := s.act(); act == 0 {
		atomic.StoreUint32(&s.idx, 1)
	} else {
		atomic.StoreUint32(&s.idx, 0)
	}
}

func (s *sched) getBuf() []schedRule {
	s.lock()
	defer s.unlock()
	return s.buf[s.act()]
}

func (s *sched) lock() {
	if atomic.LoadUint32(&s.flag) == 1 {
		s.mux.Lock()
	}
}

func (s *sched) unlock() {
	if atomic.LoadUint32(&s.flag) == 1 {
		s.mux.Unlock()
	}
}
