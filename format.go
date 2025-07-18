package clock

import (
	"strconv"
	"time"

	"github.com/koykov/bytealg"
	"github.com/koykov/byteconv"
)

const (
	Layout      = "%m/%d %H:%M:%S%p '%y %z"
	ANSIC       = "%a %b %e %H:%M:%S %Y"
	UnixDate    = "%a %b %e %H:%M:%S %Z %Y"
	RubyDate    = "%a %b %d %H:%M:%S %z %Y"
	RFC822      = "%d %b %y %H:%M %Z"
	RFC822Z     = "%d %b %y %H:%M %z"
	RFC850      = "%A, %d-%b-%y %H:%M:%S %Z"
	RFC1123     = "%a, %d %b %Y %H:%M:%S %Z"
	RFC1123Z    = "%a, %d %b %Y %H:%M:%S %z"
	RFC3339     = "%Y-%m-%dT%H:%M:%S%z"
	RFC3339Nano = "%Y-%m-%dT%H:%M:%S.%n%z"
	Kitchen     = "%L:%M%p"
	Stamp       = "%b %e %H:%M:%S"
	StampMilli  = "%b %e %H:%M:%S.%i"
	StampMicro  = "%b %e %H:%M:%S.%o"
	StampNano   = "%b %e %H:%M:%S.%N"
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

// DEPRECATED: use FormatString instead.
func FormatStr(format string, datetime time.Time) (string, error) {
	return FormatString(format, datetime)
}

func FormatString(format string, datetime time.Time) (string, error) {
	r, err := AppendFormat(nil, format, datetime)
	if err != nil {
		return "", err
	}
	return byteconv.B2S(r), nil
}

func appendFmt(buf []byte, format string, t time.Time) ([]byte, error) {
	off := 0
	for {
		p := bytealg.IndexAtString(format, "%", off)
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
		case 'b', 'h':
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
		case 'L':
			hour := t.Hour()
			buf = strconv.AppendInt(buf, int64(hour%12), 10)
		case 'M':
			mn := t.Minute()
			buf = appendInt(buf, mn, 2, '0')
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
		case 'i':
			ns := t.Unix()*1e3 + int64(t.Nanosecond())/1e6
			buf = appendInt(buf, int(ns%1e3), 3, '0')
		case 'o':
			us := t.Unix()*1e6 + int64(t.Nanosecond())/1e3
			buf = appendInt(buf, int(us%1e6), 6, '0')
		case 'n':
			ns := t.Nanosecond()
			buf = appendNano(buf, ns, 7)
		case 'N':
			ns := t.Nanosecond()
			buf = appendInt(buf, ns, 9, '0')
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
		case 'v':
			buf, _ = appendFmt(buf, "%e-%b-%Y", t)
		case 'x':
			buf, _ = appendFmt(buf, "%m/%d/%y", t)
		// timezones
		case 'z':
			_, offset := t.Zone()
			if offset == 0 {
				buf = append(buf, 'Z')
			} else {
				if offset < 0 {
					buf = append(buf, '-')
					offset = -offset
				} else {
					buf = append(buf, '+')
				}
				buf = appendInt(buf, offset/3600, 2, '0')
				// buf = append(buf, ':')
				buf = appendInt(buf, offset/60%60, 2, '0')
			}
		case 'Z':
			zone, _ := t.Zone()
			buf = append(buf, zone...)
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

func appendNano(buf []byte, x, w int) []byte {
	for x > 0 {
		if x%10 > 0 {
			break
		}
		x = x / 10
	}
	return appendInt(buf, x, w, '0')
}
