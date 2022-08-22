package clock

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Run("layout", func(t *testing.T) {
		t.Log(time.Now().Format(time.Layout))
	})
	t.Run("ansic", func(t *testing.T) {
		t.Log(time.Now().Format(time.ANSIC))
	})
	t.Run("unixdate", func(t *testing.T) {
		t.Log(time.Now().Format(time.UnixDate))
	})
	t.Run("rubydate", func(t *testing.T) {
		t.Log(time.Now().Format(time.RubyDate))
	})
	t.Run("RFC822", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC822))
	})
	t.Run("RFC822Z", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC822Z))
	})
	t.Run("RFC850", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC850))
	})
	t.Run("RFC1123", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC1123))
	})
	t.Run("RFC1123Z", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC1123Z))
	})
	t.Run("RFC3339", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC3339))
	})
	t.Run("RFC3339Nano", func(t *testing.T) {
		t.Log(time.Now().Format(time.RFC3339Nano))
	})
	t.Run("Kitchen", func(t *testing.T) {
		t.Log(time.Now().Format(time.Kitchen))
	})
	t.Run("Stamp", func(t *testing.T) {
		t.Log(time.Now().Format(time.Stamp))
	})
	t.Run("StampMilli", func(t *testing.T) {
		t.Log(time.Now().Format(time.StampMilli))
	})
	t.Run("StampMicro", func(t *testing.T) {
		t.Log(time.Now().Format(time.StampMicro))
	})
	t.Run("StampNano", func(t *testing.T) {
		t.Log(time.Now().Format(time.StampNano))
	})
}
