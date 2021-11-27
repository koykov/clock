package clock

import (
	"testing"
	"time"
)

func TestClock(t *testing.T) {
	t.Run("diff time.Nov() - clock.Now()", func(t *testing.T) {
		const (
			prec  = time.Millisecond
			allow = prec + prec/2
		)
		var (
			diff time.Duration
			fail bool
		)
		c := NewClockWP(prec)
		c.Start()
		for i := 0; i < 500; i++ {
			if diff = time.Now().Sub(c.Now()); diff > allow {
				fail = true
				break
			}
			time.Sleep(time.Millisecond)
		}
		c.Stop()
		if fail {
			t.Errorf("diff is too big: %s, max %s", diff, allow)
		}
	})
	t.Run("jump", func(t *testing.T) {
		c := NewClock()
		c.Start()
		n0 := c.Now()
		c.Jump(time.Minute * 2)
		c.Jump(-time.Minute)
		n1 := c.Now()
		if diff := n1.Sub(n0); diff < time.Minute-time.Millisecond || diff > time.Minute+time.Millisecond {
			t.Errorf("diff is too big: %s, max %s", diff, time.Minute)
		}
		c.Stop()
	})
}

func BenchmarkClock(b *testing.B) {
	b.Run("clock.Now()", func(b *testing.B) {
		c := NewClock()
		c.Start()
		var n time.Time
		for i := 0; i < b.N; i++ {
			n = c.Now()
		}
		_ = n
		c.Stop()
	})
	b.Run("time.Now()", func(b *testing.B) {
		var n time.Time
		for i := 0; i < b.N; i++ {
			n = time.Now()
		}
		_ = n
	})
}
