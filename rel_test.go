package clock

import (
	"testing"
	"time"
)

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
