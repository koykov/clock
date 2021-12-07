package clock

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		var (
			a  uint32
			fn = func() {
				atomic.AddUint32(&a, 1)
			}
		)
		c := NewClock()
		c.Start()
		c.Schedule(time.Millisecond*5, fn)
		time.Sleep(time.Millisecond)
		if v := atomic.LoadUint32(&a); v != 0 {
			t.Errorf("wrong value: need %d, got %d", 0, v)
		}
		time.Sleep(time.Millisecond * 5)
		if v := atomic.LoadUint32(&a); v != 1 {
			t.Errorf("wrong value: need %d, got %d", 1, v)
		}
		c.Stop()
	})
	t.Run("jump", func(t *testing.T) {
		var (
			a  uint32
			fn = func() {
				atomic.AddUint32(&a, 1)
			}
		)
		c := NewClock()
		c.Start()
		c.Schedule(time.Millisecond*5, fn)
		c.Jump(time.Millisecond * 5)
		time.Sleep(time.Millisecond)
		if v := atomic.LoadUint32(&a); v != 1 {
			t.Errorf("wrong value: need %d, got %d", 1, v)
		}
		c.Stop()
	})
}
