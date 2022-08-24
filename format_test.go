package clock

import (
	"bytes"
	"testing"
	"time"

	"github.com/koykov/fastconv"
)

func TestFormat(t *testing.T) {
	assert := func(t *testing.T, dt time.Time, format, expect string, err error) {
		r, err1 := FormatStr(format, dt)
		if err != nil && err1 != err {
			t.FailNow()
			return
		}
		if r != expect {
			t.FailNow()
		}
	}
	now, _ := time.Parse("2006-01-02", "1997-04-19")
	t.Run("eof", func(t *testing.T) {
		assert(t, now, "unexpected EOF: %", "", ErrBadEOF)
	})
	t.Run("mod", func(t *testing.T) {
		assert(t, now, "mod symbol: %%", "mod symbol: %", nil)
	})

	t.Run("year short", func(t *testing.T) {
		assert(t, now, "year short: %y", "year short: 97", nil)
	})
	t.Run("year", func(t *testing.T) {
		assert(t, now, "year short: %Y", "year short: 1997", nil)
	})
	t.Run("century", func(t *testing.T) {
		assert(t, now, "year short: %C", "year short: 19", nil)
	})

	t.Run("month", func(t *testing.T) {
		assert(t, now, "month: %m", "month: 04", nil)
	})
	t.Run("month short", func(t *testing.T) {
		assert(t, now, "month short: %b", "month short: Apr", nil)
	})
	t.Run("month long", func(t *testing.T) {
		assert(t, now, "month long: %B", "month long: April", nil)
	})

	t.Run("week number (sun)", func(t *testing.T) {
		assert(t, now, "week number (sun): %U", "week number (sun): 15", nil)
	})
	t.Run("week number (iso)", func(t *testing.T) {
		assert(t, now, "week number (iso): %V", "week number (iso): 16", nil)
	})
	t.Run("week number (mon)", func(t *testing.T) {
		assert(t, now, "week number (mon): %W", "week number (mon): 15", nil)
	})

	t.Run("day", func(t *testing.T) {
		assert(t, now, "day: %d", "day: 19", nil)
	})
	t.Run("day of month (space)", func(t *testing.T) {
		assert(t, now, "day: %e", "day: 19", nil)
	})
	t.Run("day of year", func(t *testing.T) {
		assert(t, now, "day of year: %j", "day of year: 109", nil)
	})
	t.Run("day of week", func(t *testing.T) {
		assert(t, now, "day of week: %w", "day of week: 6", nil)
	})
	t.Run("day of week (iso)", func(t *testing.T) {
		assert(t, now, "day of week (iso): %u", "day of week (iso): 6", nil)
	})
	t.Run("day short", func(t *testing.T) {
		assert(t, now, "day short: %a", "day short: Sat", nil)
	})
	t.Run("day long", func(t *testing.T) {
		assert(t, now, "day long: %A", "day long: Saturday", nil)
	})

	dt := time.Unix(1136239445, 123456789).UTC()
	t.Run("hour (zero pad)", func(t *testing.T) {
		assert(t, dt, "hour (zero pad): %H", "hour (zero pad): 22", nil)
	})
	t.Run("hour (space pad)", func(t *testing.T) {
		assert(t, dt, "hour (space pad): %k", "hour (space pad): 22", nil)
	})
	t.Run("hour12 (zero pad)", func(t *testing.T) {
		assert(t, dt, "hour12 (zero pad): %I", "hour12 (zero pad): 10", nil)
	})
	t.Run("hour12 (space pad)", func(t *testing.T) {
		assert(t, dt, "hour12 (space pad): %l", "hour12 (space pad): 10", nil)
	})
	t.Run("minute", func(t *testing.T) {
		assert(t, dt, "minute: %M", "minute: 04", nil)
	})
	t.Run("second", func(t *testing.T) {
		assert(t, dt, "second: %S", "second: 05", nil)
	})
	t.Run("AM/PM", func(t *testing.T) {
		assert(t, dt, "AM/PM: %p", "AM/PM: PM", nil)
	})
	t.Run("am/pm", func(t *testing.T) {
		assert(t, dt, "am/pm: %P", "am/pm: pm", nil)
	})
	t.Run("nat time", func(t *testing.T) {
		assert(t, dt, "nat time: %X", "nat time: 22:04:05", nil)
	})
	t.Run("complex r", func(t *testing.T) {
		assert(t, dt, "complex r: %r", "complex r: 10:04:05 PM", nil)
	})
	t.Run("complex R", func(t *testing.T) {
		assert(t, dt, "complex R: %r", "complex R: 10:04", nil)
	})
	t.Run("complex T", func(t *testing.T) {
		assert(t, dt, "complex R: %r", "complex R: 10:04:05", nil)
	})
}

func TestFormatNativeLayout(t *testing.T) {
	t.Run("Layout", func(t *testing.T) {
		t.Log(time.Now().Format(time.Layout))
	})
	t.Run("ANSIC", func(t *testing.T) {
		t.Log(time.Now().Format(time.ANSIC))
	})
	t.Run("UnixDate", func(t *testing.T) {
		t.Log(time.Now().Format(time.UnixDate))
	})
	t.Run("RubyDate", func(t *testing.T) {
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

func TestFormatInternal(t *testing.T) {
	assert := func(t *testing.T, buf []byte, x, w int, pad byte, expect []byte) {
		buf = buf[:0]
		buf = appendInt(buf, x, w, pad)
		if !bytes.Equal(buf, expect) {
			t.FailNow()
		}
	}
	var buf []byte
	t.Run("appendInt 2018 2", func(t *testing.T) {
		assert(t, buf, 2018, 2, '0', []byte("18"))
	})
	t.Run("appendInt 1997 4", func(t *testing.T) {
		assert(t, buf, 1997, 4, '0', []byte("1997"))
	})
	t.Run("appendInt 34 4", func(t *testing.T) {
		assert(t, buf, 34, 4, '0', []byte("0034"))
	})
}

func BenchmarkFormat(b *testing.B) {
	assert := func(b *testing.B, dt time.Time, format, expect string, err error) {
		var (
			buf  []byte
			err1 error
		)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf, err1 = AppendFormat(buf, format, dt)
			if err != nil && err1 != err {
				b.FailNow()
				return
			}
			if fastconv.B2S(buf) != expect {
				b.FailNow()
			}
		}
	}
	now, _ := time.Parse("2006", "1997")
	b.Run("eof", func(b *testing.B) {
		assert(b, now, "unexpected EOF: %", "", ErrBadEOF)
	})
	b.Run("mod", func(b *testing.B) {
		assert(b, now, "mod symbol: %%", "mod symbol: %", nil)
	})
	b.Run("year short", func(b *testing.B) {
		assert(b, now, "year short: %y", "year short: 97", nil)
	})
	b.Run("year", func(b *testing.B) {
		assert(b, now, "year short: %Y", "year short: 1997", nil)
	})
	b.Run("century", func(b *testing.B) {
		assert(b, now, "year short: %C", "year short: 19", nil)
	})
}

func BenchmarkFormatInternal(b *testing.B) {
	assert := func(b *testing.B, buf []byte, x, w int, pad byte, expect []byte) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = appendInt(buf, x, w, pad)
			if !bytes.Equal(buf, expect) {
				b.FailNow()
			}
		}
	}
	var buf []byte
	b.Run("appendInt 2018 2", func(b *testing.B) {
		assert(b, buf, 2018, 2, '0', []byte("18"))
	})
	b.Run("appendInt 1997 4", func(b *testing.B) {
		assert(b, buf, 1997, 4, '0', []byte("1997"))
	})
	b.Run("appendInt 34 4", func(b *testing.B) {
		assert(b, buf, 34, 4, '0', []byte("0034"))
	})
}
