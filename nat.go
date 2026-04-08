package clock

import "time"

var natNum = map[string]int{
	"0":         0,
	"1":         1,
	"2":         2,
	"3":         3,
	"4":         4,
	"5":         5,
	"6":         6,
	"7":         7,
	"8":         8,
	"9":         9,
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighteen":  18,
	"nineteen":  19,
}
var natUnit = map[string]time.Duration{
	"now": 0,
	// ...
}

func RelativeNatural(raw string) (dur time.Duration, err error) {
	var num int
	var d time.Duration
	for {
		raw1, tkn, ok := natToken(raw)
		raw = raw1
		if !ok {
			return 0, nil
		}
		num1, numOK := natNum[tkn]
		dur1, durOK := natUnit[tkn]
		switch {
		case numOK && durOK:
			num, d = num1, dur1
			return -time.Duration(num) * d, nil
		case numOK && !durOK:
			// ???
		case !numOK && durOK:
			num, d = 1, dur
			return -time.Duration(num) * d, nil
		case !numOK && !durOK:
			continue
		}
	}
}

func natToken(s string) (string, string, bool) {
	s = skipWS(s)
	if len(s) == 0 {
		return s, "", false
	}
	for i := 0; i < len(s); i++ {
		if c := s[i]; c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			return s[i:], s[:i], true
		}
	}
	return "", s, true
}

func skipWS(s string) string {
	var off int
	for i := 0; i < len(s); i++ {
		if c := s[i]; c != ' ' && c != '\n' && c != '\r' && c != '\t' {
			break
		}
		off++
	}
	return s[off:]
}
