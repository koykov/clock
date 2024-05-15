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
		for i := 0; i < 10; i++ {
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

func TestRelative(t *testing.T) {
	spans := []struct {
		key, exp string
	}{
		{"2h 30min", "2h30m0s"},
		{"2 h", "2h0m0s"},
		{"2hours", "2h0m0s"},
		{"48hr", "48h0m0s"},
		{"1y 12month", "17532h43m12s"},
		{"55s500ms", "55.5s"},
		{"300ms20s 5day", "120h0m20.3s"},
		{"-2h 30min", "-2h30m0s"},
		{"-2 h", "-2h0m0s"},
		{"-2hours", "-2h0m0s"},
		{"-48hr", "-48h0m0s"},
		{"-1y 12month", "-17532h43m12s"},
		{"-55s500ms", "-55.5s"},
		{"-300ms20s 5day", "-120h0m20.3s"},
		{"2century 43 y 3M 3 w 15d  17 h 43m  34 s 400ms 123 us  55 ns", "2132850h41m22.400000055s"},
	}
	for _, span := range spans {
		t.Run(span.key, func(t *testing.T) {
			c := NewClock()
			c.Start()
			if ts := c.Relative(span.key).Sub(c.Now()).String(); ts != span.exp {
				t.Errorf("relative fail: need %s, got %s", span.exp, ts)
			}
			c.Stop()
		})
	}
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

func BenchmarkRelative(b *testing.B) {
	span, exp := "300ms20s 5day", 120*time.Hour+20*time.Second+300*time.Millisecond
	c := NewClockWP(time.Second)
	c.Start()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if ts := c.Relative(span).Sub(c.Now()); ts != exp {
			b.Errorf("relative fail: need %s, got %s", exp, ts)
		}
	}
	c.Stop()
}
