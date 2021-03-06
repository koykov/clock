package clock

import (
	"time"
)

// Stuck represents stuck clock implements Interface.
type Stuck struct {
	sec, nsec int64
}

func NewStuck(sec, nsec int64) *Stuck {
	c := &Stuck{
		sec:  sec,
		nsec: nsec,
	}
	return c
}

func (c Stuck) Now() time.Time {
	return time.Unix(c.sec, c.nsec)
}
