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

var (
	longDayNames = []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	shortDayNames = []string{
		"Sun",
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
	}
	shortMonthNames = []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}
	longMonthNames = []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
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
		if p == -1 {
			buf = append(buf, format[off:]...)
			return buf, nil
		}
		if p == len(format)-1 {
			return buf, ErrBadEOF
		}
		if p-1 > off {
			buf = append(buf, format[off:p]...)
		}
		switch format[p+1] {
		case '%':
			buf = append(buf, '%')
		// year
		case 'y':
			year := dt.Year()
			buf = appendInt(buf, year%100, 2)
		case 'Y':
			year := dt.Year()
			buf = appendInt(buf, year, 4)
		case 'C':
			year := dt.Year()
			buf = strconv.AppendInt(buf, int64(year/100), 10)
		// month
		case 'm':
			month := dt.Month()
			buf = appendInt(buf, int(month), 2)
		case 'b':
			month := dt.Month()
			buf = append(buf, shortMonthNames[month-1]...)
		case 'B':
			month := dt.Month()
			buf = append(buf, longMonthNames[month-1]...)
		// week
		case 'U':
			yd := dt.YearDay()
			wd := int(dt.Weekday())
			if yd < wd {
				buf = append(buf, '0', '0')
				return buf, nil
			}
			n := ((yd - wd) / 7) + 1
			buf = appendInt(buf, n, 2)
		case 'V':
			_, w := dt.ISOWeek()
			buf = appendInt(buf, w, 2)
		case 'W':
			yd := dt.YearDay()
			wd := int(dt.Weekday())
			off1 := wd - 1
			if off1 < 0 {
				off1 += 7
			}
			if yd < off1 {
				buf = append(buf, '0', '0')
				return buf, nil
			}
			n := ((yd - off1) / 7) + 1
			buf = appendInt(buf, n, 2)
		// day
		case 'd':
			day := dt.Day()
			buf = appendInt(buf, day, 2)
		case 'j':
			day := dt.YearDay()
			buf = appendInt(buf, day, 3)
		case 'w':
			day := dt.Weekday()
			buf = append(buf, byte('0'+int(day)))
		case 'u':
			day := dt.Weekday()
			if day < 1 {
				day += 7
			}
			buf = append(buf, byte('0'+day))
		case 'a':
			day := dt.Weekday()
			buf = append(buf, shortDayNames[day]...)
		case 'A':
			day := dt.Weekday()
			buf = append(buf, longDayNames[day]...)
		case 'e':
			day := dt.Day()
			if day < 10 {
				buf = append(buf, ' ')
			}
			buf = strconv.AppendInt(buf, int64(day), 10)
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
