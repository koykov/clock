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

	_ = Format
)

func AppendFormat(dst []byte, format string, datetime time.Time) ([]byte, error) {
	return appendFmt(dst, format, datetime)
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

func appendFmt(buf []byte, format string, t time.Time) ([]byte, error) {
	off := 0
	for {
		p := bytealg.IndexAtStr(format, "%", off)
		if p == -1 {
			buf = append(buf, format[off:]...)
			return buf, nil
		}
		if p == len(format)-1 {
			return buf, ErrBadEOF
		}
		if p-1 >= off {
			buf = append(buf, format[off:p]...)
		}
		switch format[p+1] {
		case '%':
			buf = append(buf, '%')
		// year
		case 'y':
			year := t.Year()
			buf = appendInt(buf, year%100, 2, '0')
		case 'Y':
			year := t.Year()
			buf = appendInt(buf, year, 4, '0')
		case 'C':
			year := t.Year()
			buf = strconv.AppendInt(buf, int64(year/100), 10)
		// month
		case 'm':
			month := t.Month()
			buf = appendInt(buf, int(month), 2, '0')
		case 'b':
			month := t.Month()
			buf = append(buf, shortMonthNames[month-1]...)
		case 'B':
			month := t.Month()
			buf = append(buf, longMonthNames[month-1]...)
		// week
		case 'U':
			yd := t.YearDay()
			wd := int(t.Weekday())
			if yd < wd {
				buf = append(buf, '0', '0')
				return buf, nil
			}
			n := ((yd - wd) / 7) + 1
			buf = appendInt(buf, n, 2, '0')
		case 'V':
			_, w := t.ISOWeek()
			buf = appendInt(buf, w, 2, '0')
		case 'W':
			yd := t.YearDay()
			wd := int(t.Weekday())
			off1 := wd - 1
			if off1 < 0 {
				off1 += 7
			}
			if yd < off1 {
				buf = append(buf, '0')
				buf = append(buf, '0')
			} else {
				n := ((yd - off1) / 7) + 1
				buf = appendInt(buf, n, 2, '0')
			}
		// day
		case 'd':
			day := t.Day()
			buf = appendInt(buf, day, 2, '0')
		case 'j':
			day := t.YearDay()
			buf = appendInt(buf, day, 3, '0')
		case 'w':
			day := t.Weekday()
			buf = append(buf, byte('0'+int(day)))
		case 'u':
			day := t.Weekday()
			if day < 1 {
				day += 7
			}
			buf = append(buf, byte('0'+day))
		case 'a':
			day := t.Weekday()
			buf = append(buf, shortDayNames[day]...)
		case 'A':
			day := t.Weekday()
			buf = append(buf, longDayNames[day]...)
		case 'e':
			day := t.Day()
			buf = appendInt(buf, day, 2, ' ')
		// time
		case 'H':
			hour := t.Hour()
			buf = appendInt(buf, hour, 2, '0')
		case 'k':
			hour := t.Hour()
			buf = appendInt(buf, hour, 2, ' ')
		case 'I':
			hour := t.Hour()
			buf = appendInt(buf, hour%12, 2, '0')
		case 'l':
			hour := t.Hour()
			buf = appendInt(buf, hour%12, 2, ' ')
		case 'M':
			min := t.Minute()
			buf = appendInt(buf, min, 2, '0')
		case 'S':
			sec := t.Second()
			buf = appendInt(buf, sec, 2, '0')
		case 'p':
			if t.Hour() > 12 {
				buf = append(buf, "PM"...)
			} else {
				buf = append(buf, "AM"...)
			}
		case 'P':
			if t.Hour() > 12 {
				buf = append(buf, "pm"...)
			} else {
				buf = append(buf, "am"...)
			}
		case 'X':
			buf, _ = appendFmt(buf, "%H:%M:%S", t)
		// complex
		case 'r':
			buf, _ = appendFmt(buf, "%I:%M:%S %p", t)
		case 'R':
			buf, _ = appendFmt(buf, "%H:%M", t)
		case 'T':
			buf, _ = appendFmt(buf, "%H:%M:%S", t)
		case 'c':
			buf, _ = appendFmt(buf, "%a %b %e %H:%M:%S %Y", t)
		case 'D':
			buf, _ = appendFmt(buf, "%m/%d/%y", t)
		case 'F':
			buf, _ = appendFmt(buf, "%Y-%m-%d", t)
		case 's':
			buf = strconv.AppendInt(buf, t.Unix(), 10)
		case 'x':
			buf = t.AppendFormat(buf, "01/02/06")
		}
		off = p + 2
	}
}

func appendInt(buf []byte, x, w int, pad byte) []byte {
	if x < 0 {
		buf = append(buf, '-')
	}
	off := len(buf)
	buf = bytealg.GrowDelta(buf, w)
	c, pad1 := w-1, false
	for c = w - 1; c >= 0; c-- {
		if pad1 {
			buf[off+c] = pad
		} else {
			buf[off+c] = byte('0' + x%10)
		}
		if x = x / 10; x == 0 {
			pad1 = true
		}
	}
	return buf
}
