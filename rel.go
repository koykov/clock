package clock

import (
	"strconv"
	"time"

	"github.com/koykov/bytealg"
)

func Relative(raw string) (dur time.Duration, err error) {
	var (
		neg bool
		off int
	)
	if raw = bytealg.TrimString(raw, " "); len(raw) == 0 {
		err = ErrNoDur
		return
	}
	if neg = raw[0] == '-'; neg {
		off++
	}
	for off < len(raw) {
		var (
			n    time.Duration
			unit string
			ok   bool
		)
		if n, off, ok = relNum(raw, off); !ok || n == 0 {
			err = ErrBadNum
			return
		}
		if unit, off = relUnit(raw, off); len(unit) == 0 {
			err = ErrBadUnit
			return
		}
		switch unit {
		case "nsec", "ns":
			dur += n
		case "usec", "us", "µs":
			dur += n * time.Microsecond
		case "msec", "ms":
			dur += n * time.Millisecond
		case "seconds", "second", "sec", "s":
			dur += n * time.Second
		case "minutes", "minute", "min", "m":
			dur += n * time.Minute
		case "hours", "hour", "hr", "h":
			dur += n * time.Hour
		case "days", "day", "d":
			dur += n * 24 * time.Hour
		case "weeks", "week", "w":
			dur += n * 168 * time.Hour
		case "months", "month", "M":
			d := 24 * time.Hour
			dur += n * (d*30 + d*44/100)
		case "years", "year", "y":
			d := 24 * time.Hour
			dur += n * (d*365 + d*1/4)
		case "century", "cen", "c":
			d := 24 * time.Hour
			dur += n * (d * 36525)
		case "millennium", "mil":
			d := 24 * time.Hour
			dur += n * (d * 365250)
		}
	}
	if neg {
		dur = -dur
	}
	return
}

func relNum(raw string, off int) (time.Duration, int, bool) {
	pos := off
loop:
	c := raw[pos]
	if c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
		pos++
		goto loop
	}
	if pos > off {
		if i, err := strconv.Atoi(raw[off:pos]); err == nil {
			return time.Duration(i), pos, true
		}
	}
	return 0, off, false
}

func relUnit(raw string, off int) (string, int) {
	if raw[off] == ' ' {
		off++
	}
	pos := off
loop:
	if pos == len(raw)-1 {
		pos++
		return raw[off:pos], pos
	}
	c := raw[pos]
	if c != '0' && c != '1' && c != '2' && c != '3' && c != '4' && c != '5' && c != '6' && c != '7' && c != '8' && c != '9' {
		pos++
		goto loop
	}
	if raw[pos-1] == ' ' {
		return raw[off : pos-1], pos
	}
	return raw[off:pos], pos
}
