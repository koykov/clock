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
	now, _ := time.Parse("2006", "1997")
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
	assert := func(t *testing.T, buf []byte, x, w int, expect []byte) {
		buf = buf[:0]
		buf = appendInt(buf, x, w)
		if !bytes.Equal(buf, expect) {
			t.FailNow()
		}
	}
	var buf []byte
	t.Run("appendInt 2018 2", func(t *testing.T) {
		assert(t, buf, 2018, 2, []byte("18"))
	})
	t.Run("appendInt 1997 4", func(t *testing.T) {
		assert(t, buf, 1997, 4, []byte("1997"))
	})
	t.Run("appendInt 34 4", func(t *testing.T) {
		assert(t, buf, 34, 4, []byte("0034"))
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
	assert := func(b *testing.B, buf []byte, x, w int, expect []byte) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = appendInt(buf, x, w)
			if !bytes.Equal(buf, expect) {
				b.FailNow()
			}
		}
	}
	var buf []byte
	b.Run("appendInt 2018 2", func(b *testing.B) {
		assert(b, buf, 2018, 2, []byte("18"))
	})
	b.Run("appendInt 1997 4", func(b *testing.B) {
		assert(b, buf, 1997, 4, []byte("1997"))
	})
	b.Run("appendInt 34 4", func(b *testing.B) {
		assert(b, buf, 34, 4, []byte("0034"))
	})
}
