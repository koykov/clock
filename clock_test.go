package clock

import (
	"testing"
	"time"
)

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
