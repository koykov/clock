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

func (c Stuck) Relative(raw string) time.Time {
	if dur, err := Relative(raw); err == nil {
		return c.Now().Add(dur)
	}
	return time.Time{}
}

func (c Stuck) RelativeNatural(raw string) time.Time {
	if dur, err := RelativeNatural(raw); err == nil {
		return c.Now().Add(dur)
	}
	return time.Time{}
}
