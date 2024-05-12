package clock

import (
	"time"

	"github.com/koykov/x2bytes"
)

// TF represents Time-Format pair.
type TF struct {
	Time   time.Time
	Format string
}

// TimeToBytes converts from time.
func TimeToBytes(dst []byte, val any) ([]byte, error) {
	format := Layout
	var t time.Time
	switch x := val.(type) {
	case time.Time:
		t = x
	case *time.Time:
		t = *x
	case TF:
		t, format = x.Time, x.Format
	case *TF:
		t, format = x.Time, x.Format
	default:
		return dst, x2bytes.ErrUnknownType
	}
	var err error
	dst, err = AppendFormat(dst, format, t)
	return dst, err
}

func init() {
	x2bytes.RegisterToBytesFn(TimeToBytes)
}
