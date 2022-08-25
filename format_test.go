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
		{
			key:    "unexpected eof",
			format: "%",
			err:    ErrBadEOF,
		},
		{
			key:    "percent",
			format: "%%",
			expect: "%",
			time:   t97,
		},
		{
			key:    "year short",
			format: "%y",
			expect: "97",
			time:   t97,
		},
		{
			key:    "year",
			format: "%Y",
			expect: "1997",
			time:   t97,
		},
		{
			key:    "century",
			format: "%C",
			expect: "19",
			time:   t97,
		},
		{
			key:    "month",
			format: "%m",
			expect: "04",
			time:   t97,
		},
		{
			key:    "month name short",
			format: "%b",
			expect: "Apr",
			time:   t97,
		},
		{
			key:    "month name",
			format: "%B",
			expect: "April",
			time:   t97,
		},
		{
			key:    "week number (sunday)",
			format: "%U",
			expect: "15",
			time:   t97,
		},
		{
			key:    "week number (monday)",
			format: "%W",
			expect: "15",
			time:   t97,
		},
		{
			key:    "week number (iso)",
			format: "%V",
			expect: "16",
			time:   t97,
		},
		{
			key:    "day",
			format: "%d",
			expect: "19",
			time:   t97,
		},
		{
			key:    "day (space pad)",
			format: "%e",
			expect: "19",
			time:   t97,
		},
		{
			key:    "day of year",
			format: "%j",
			expect: "109",
			time:   t97,
		},
		{
			key:    "day of week",
			format: "%w",
			expect: "6",
			time:   t97,
		},
		{
			key:    "day of week (iso)",
			format: "%u",
			expect: "6",
			time:   t97,
		},
		{
			key:    "day name short",
			format: "%a",
			expect: "Sat",
			time:   t97,
		},
		{
			key:    "day name",
			format: "%A",
			expect: "Saturday",
			time:   t97,
		},
		{
			key:    "hour",
			format: "%H",
			expect: "22",
			time:   t0,
		},
		{
			key:    "hour (space pad)",
			format: "%k",
			expect: "22",
			time:   t0,
		},
		{
			key:    "hour 12",
			format: "%I",
			expect: "10",
			time:   t0,
		},
		{
			key:    "hour 12 (space pad)",
			format: "%l",
			expect: "10",
			time:   t0,
		},
		{
			key:    "minute",
			format: "%M",
			expect: "04",
			time:   t0,
		},
		{
			key:    "second",
			format: "%S",
			expect: "05",
			time:   t0,
		},
		{
			key:    "AM/PM",
			format: "%p",
			expect: "PM",
			time:   t0,
		},
		{
			key:    "am/pm",
			format: "%P",
			expect: "pm",
			time:   t0,
		},
		{
			key:    "preferred time",
			format: "%X",
			expect: "22:04:05",
			time:   t0,
		},
		{
			key:    "unixtime",
			format: "%s",
			expect: "1136239445",
			time:   t0,
		},
		{
			key:    "complex/r",
			format: "%r",
			expect: "10:04:05 PM",
			time:   t0,
		},
		{
			key:    "complex/R",
			format: "%R",
			expect: "22:04",
			time:   t0,
		},
		{
			key:    "complex/T",
			format: "%T",
			expect: "22:04:05",
			time:   t0,
		},
		{
			key:    "complex/c",
			format: "%c",
			expect: "Mon Jan  2 22:04:05 2006",
			time:   t0,
		},
		{
			key:    "complex/D",
			format: "%D",
			expect: "04/19/97",
			time:   t97,
		},
		{
			key:    "complex/F",
			format: "%F",
			expect: "1997-04-19",
			time:   t97,
		},
	}

	stagesRFC3339 = []stageRFC3339{
		{
			time.Date(2008, 9, 17, 20, 4, 26, 0, time.UTC),
			"2008-09-17T20:04:26Z",
		},
		{
			time.Date(1994, 9, 17, 20, 4, 26, 0, time.FixedZone("EST", -18000)),
			"1994-09-17T20:04:26-05:00",
		},
		{
			time.Date(2000, 12, 26, 1, 15, 6, 0, time.FixedZone("OTO", 15600)),
			"2000-12-26T01:15:06+04:20",
		},
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
	for _, f := range stagesRFC3339 {
		r, _ := FormatStr(RFC3339, f.time)
		if r != f.expect {
			t.Errorf("format RFC3339 mismatch: '%s' vs '%s'", r, f.expect)
		}
	}
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

func BenchmarkFormat(b *testing.B) {
	for _, stage := range stagesFmt {
		b.Run(stage.key, func(b *testing.B) {
			var (
				buf []byte
				err error
			)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf = buf[:0]
				buf, err = AppendFormat(buf, stage.format, stage.time)
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
