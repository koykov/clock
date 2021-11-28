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
	if raw = bytealg.TrimStr(raw, " "); len(raw) == 0 {
		err = ErrNoDur
		return
	}
	raw = bytealg.ToLowerStr(raw)
	if neg = raw[0] == '-'; neg {
		off++
	}
	for off < len(raw) {
		var (
			n    int
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
		// todo use n and unit
	}
	return
}

func relNum(raw string, off int) (int, int, bool) {
	pos := off
loop:
	c := raw[pos]
	if c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
		pos++
		goto loop
	}
	if pos > off {
		if i, err := strconv.Atoi(raw[off:pos]); err == nil {
			return i, pos, true
		}
	}
	return 0, off, false
}

func relUnit(raw string, off int) (string, int) {
	pos := off
loop:
	if pos == len(raw)-1 {
		return raw[off:pos], pos
	}
	c := raw[pos]
	if c == ' ' {
		off++
		pos++
		goto loop
	}
	if c != '0' && c != '1' && c != '2' && c != '3' && c != '4' && c != '5' && c != '6' && c != '7' && c != '8' && c != '9' {
		pos++
		goto loop
	}
	return raw[off:pos], pos
}
