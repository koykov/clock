package clock

import (
	"strconv"
	"time"

	"github.com/koykov/bytealg"
	"github.com/koykov/fastconv"
)

const (
	Layout      = "%m/%d %h:%m:%s%p '%y %z"
	ANSIC       = "%a %b %d %h:%m:%s %Y"
	UnixDate    = "%a %b %d %h:%m:%s %Z %Y"
	RubyDate    = "%a %b %d %h:%m:%s %z %Y"
	RFC822      = "%d %b %y %h:%m %Z"
	RFC822Z     = "%d %b %y %h:%m %z"
	RFC850      = "%A, %d-%b-%y %h:%m:%s %Z"
	RFC1123     = "%a, %d %b %Y %h:%m:%s %Z"
	RFC1123Z    = "%a, %d %b %Y %h:%m:%s %z"
	RFC3339     = "%Y-%m-%dT%h:%m:%dZ%o"
	RFC3339Nano = "%Y-%m-%dT%h:%m:%d.%nZ%o"
	Kitchen     = "%h:%m%p"
	Stamp       = "%b %d %h:%m:%s"
	StampMilli  = "b %d %h:%m:%s.%i"
	StampMicro  = "b %d %h:%m:%s.%u"
	StampNano   = "b %d %h:%m:%s.%n"
)

func AppendFormat(dst []byte, format string, datetime time.Time) ([]byte, error) {
	return appendFmt(dst, fastconv.S2B(format), datetime)
}

func Format(format string, datetime time.Time) ([]byte, error) {
	return AppendFormat(nil, format, datetime)
}

func FormatStr(format string, datetime time.Time) (string, error) {
	r, err := AppendFormat(nil, format, datetime)
	if err != nil {
		return "", err
	}
	return fastconv.B2S(r), nil
}

func appendFmt(buf []byte, format []byte, dt time.Time) ([]byte, error) {
	off := 0
	for {
		p := bytealg.IndexByteAtLR(format, '%', off)
		if p == -1 || p == len(format)-1 {
			buf = append(buf, format[off:]...)
			return buf, nil
		}
		if p-1 > off {
			buf = append(buf, format[off:p]...)
		}
		switch format[p+1] {
		case '%':
			buf = append(buf, '%')
		case 'y':
			year := dt.Year()
			buf = appendInt(buf, year%100, 2)
		case 'Y':
			year := dt.Year()
			buf = appendInt(buf, year, 4)
		case 'C':
			year := dt.Year()
			buf = strconv.AppendInt(buf, int64(year/100), 10)
		}
		off = p + 2
	}
}

func appendInt(buf []byte, x, w int) []byte {
	if x < 0 {
		buf = append(buf, '-')
	}
	off := len(buf)
	buf = bytealg.GrowDelta(buf, w)
	c, pad := w-1, false
	for c = w - 1; c >= 0; c-- {
		if pad {
			buf[off+c] = '0'
		} else {
			buf[off+c] = byte('0' + x%10)
		}
		if x = x / 10; x == 0 {
			pad = true
		}
	}
	return buf
}
