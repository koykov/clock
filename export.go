package clock

import (
	"strings"
	"time"

	"github.com/koykov/byteconv"
	"github.com/koykov/x2bytes"
)

// TimeToBytes converts from time.
func TimeToBytes(dst []byte, val any, args ...any) ([]byte, error) {
	format := Layout
	var t time.Time
	switch x := val.(type) {
	case time.Time:
		t = x
	case *time.Time:
		t = *x
	default:
		return dst, x2bytes.ErrUnknownType
	}
	if len(args) > 0 {
		switch x := args[0].(type) {
		case string:
			format = x
		case *string:
			format = *x
		case []byte:
			format = byteconv.B2S(x)
		case *[]byte:
			format = byteconv.B2S(*x)
		}
	}
	if strings.IndexByte(format, '%') == -1 {
		format = time.Layout
		dst = t.AppendFormat(dst, format)
		return dst, nil
	}
	var err error
	dst, err = AppendFormat(dst, format, t)
	return dst, err
}

func init() {
	x2bytes.RegisterToBytesFn(TimeToBytes)
}
