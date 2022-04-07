package clock

import (
	"testing"
	"time"
)

func TestStuck(t *testing.T) {
	c := NewStuck(1645495200, 0)
	x := c.Now().Format(time.RFC3339Nano)
	time.Sleep(time.Second)
	y := c.Now().Format(time.RFC3339Nano)
	if x != y {
		t.FailNow()
	}
}
