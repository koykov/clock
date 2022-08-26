package clock

import (
	"testing"
	"time"

	"github.com/koykov/fastconv"
)

type stageFmt struct {
	key,
	format,
	expect string
	time time.Time
	err  error
}

type stageRFC3339 struct {
	time   time.Time
	expect string
}

var (
	t97, _    = time.Parse("2006-01-02", "1997-04-19")
	t0        = time.Unix(1136239445, 123456789).UTC()
	stagesFmt = []stageFmt{
		{key: "unexpected eof", format: "%", err: ErrBadEOF},
		{key: "percent", format: "%%", expect: "%", time: t97},
		{key: "year short", format: "%y", expect: "97", time: t97},
		{key: "year", format: "%Y", expect: "1997", time: t97},
		{key: "century", format: "%C", expect: "19", time: t97},
		{key: "month", format: "%m", expect: "04", time: t97},
		{key: "month name short", format: "%b", expect: "Apr", time: t97},
		{key: "month name", format: "%B", expect: "April", time: t97},
		{key: "week number (sunday)", format: "%U", expect: "15", time: t97},
		{key: "week number (monday)", format: "%W", expect: "15", time: t97},
		{key: "week number (iso)", format: "%V", expect: "16", time: t97},
		{key: "day", format: "%d", expect: "19", time: t97},
		{key: "day (space pad)", format: "%e", expect: "19", time: t97},
		{key: "day of year", format: "%j", expect: "109", time: t97},
		{key: "day of week", format: "%w", expect: "6", time: t97},
		{key: "day of week (iso)", format: "%u", expect: "6", time: t97},
		{key: "day name short", format: "%a", expect: "Sat", time: t97},
		{key: "day name", format: "%A", expect: "Saturday", time: t97},
		{key: "hour", format: "%H", expect: "22", time: t0},
		{key: "hour (space pad)", format: "%k", expect: "22", time: t0},
		{key: "hour 12", format: "%I", expect: "10", time: t0},
		{key: "hour 12 (space pad)", format: "%l", expect: "10", time: t0},
		{key: "minute", format: "%M", expect: "04", time: t0},
		{key: "second", format: "%S", expect: "05", time: t0},
		{key: "AM/PM", format: "%p", expect: "PM", time: t0},
		{key: "am/pm", format: "%P", expect: "pm", time: t0},
		{key: "preferred time", format: "%X", expect: "22:04:05", time: t0},
		{key: "unixtime", format: "%s", expect: "1136239445", time: t0},
		{key: "complex/r", format: "%r", expect: "10:04:05 PM", time: t0},
		{key: "complex/R", format: "%R", expect: "22:04", time: t0},
		{key: "complex/T", format: "%T", expect: "22:04:05", time: t0},
		{key: "complex/c", format: "%c", expect: "Mon Jan  2 22:04:05 2006", time: t0},
		{key: "complex/D", format: "%D", expect: "04/19/97", time: t97},
		{key: "complex/F", format: "%F", expect: "1997-04-19", time: t97},
	}

	stagesRFC3339 = []stageRFC3339{
		{time.Date(2008, 9, 17, 20, 4, 26, 0, time.UTC), "2008-09-17T20:04:26Z"},
		{time.Date(1994, 9, 17, 20, 4, 26, 0, time.FixedZone("EST", -18000)), "1994-09-17T20:04:26-05:00"},
		{time.Date(2000, 12, 26, 1, 15, 6, 0, time.FixedZone("OTO", 15600)), "2000-12-26T01:15:06+04:20"},
	}

	tNative      = time.Unix(0, 1233810057012345600)
	stagesNative = []stageFmt{
		{key: "ANSIC", format: ANSIC, expect: "Thu Feb  5 07:00:57 2009", time: tNative},
		{key: "UnixDate", format: UnixDate, expect: "Thu Feb  5 07:00:57 EET 2009", time: tNative},
		{key: "RubyDate", format: RubyDate, expect: "Thu Feb 05 07:00:57 +0200 2009", time: tNative},
		{key: "RFC822", format: RFC822, expect: "05 Feb 09 07:00 EET", time: tNative},
		{key: "RFC850", format: RFC850, expect: "Thursday, 05-Feb-09 07:00:57 EET", time: tNative},
		{key: "RFC1123", format: RFC1123, expect: "Thu, 05 Feb 2009 07:00:57 EET", time: tNative},
		{key: "RFC1123Z", format: RFC1123Z, expect: "Thu, 05 Feb 2009 07:00:57 +0200", time: tNative},
		{key: "RFC3339", format: RFC3339, expect: "2009-02-05T07:00:57+02:00", time: tNative},
		{key: "RFC3339Nano", format: RFC3339Nano, expect: "2009-02-05T07:00:57.0123456+02:00", time: tNative},
		{key: "Kitchen", format: Kitchen, expect: "7:00AM", time: tNative},
		{key: "Stamp", format: Stamp, expect: "Feb  5 07:00:57", time: tNative},
		{key: "StampMilli", format: StampMilli, expect: "Feb  5 07:00:57.012", time: tNative},
		{key: "StampMicro", format: StampMicro, expect: "Feb  5 07:00:57.012345", time: tNative},
		{key: "StampNano", format: StampNano, expect: "Feb  5 07:00:57.012345600", time: tNative},
	}
)

func TestFormat(t *testing.T) {
	for _, stage := range stagesFmt {
		t.Run(stage.key, func(t *testing.T) {
			r, err := FormatStr(stage.format, stage.time)
			if stage.err != nil {
				if err != stage.err {
					t.Errorf("error mismatch: '%s' vs '%s'", err.Error(), stage.err.Error())
				}
				return
			}
			if r != stage.expect {
				t.Errorf("format mismatch: '%s' vs '%s'", r, stage.expect)
			}
		})
	}
}

func TestFormatRFC3339(t *testing.T) {
	for _, stage := range stagesRFC3339 {
		r, _ := FormatStr(RFC3339, stage.time)
		if r != stage.expect {
			t.Errorf("format RFC3339 mismatch: '%s' vs '%s'", r, stage.expect)
		}
	}
}

func TestFormatNativeLayout(t *testing.T) {
	for _, stage := range stagesNative {
		t.Run(stage.key, func(t *testing.T) {
			r, _ := FormatStr(stage.format, stage.time)
			if r != stage.expect {
				t.Errorf("format native mismatch: '%s' vs '%s'", r, stage.expect)
			}
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	for _, stage := range stagesFmt {
		b.Run(stage.key, func(b *testing.B) {
			var (
				buf []byte
				err error
			)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf, err = AppendFormat(buf[:0], stage.format, stage.time)
				if stage.err != nil {
					if err != stage.err {
						b.Errorf("error mismatch: '%s' vs '%s'", err.Error(), stage.err.Error())
					}
					return
				}
				if fastconv.B2S(buf) != stage.expect {
					b.Errorf("format mismatch: '%s' vs '%s'", string(buf), stage.expect)
				}
			}
		})
	}
}

func BenchmarkFormatNativeLayout(b *testing.B) {
	for _, stage := range stagesNative {
		b.Run(stage.key, func(b *testing.B) {
			var buf []byte
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf, _ = AppendFormat(buf[:0], stage.format, stage.time)
				if fastconv.B2S(buf) != stage.expect {
					b.Errorf("format native mismatch: '%s' vs '%s'", string(buf), stage.expect)
				}
			}
		})
	}
}
