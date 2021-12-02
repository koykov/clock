package clock

import "errors"

var (
	ErrNoDur       = errors.New("no duration passed")
	ErrBadNum      = errors.New("bad span number")
	ErrBadUnit     = errors.New("bad span unit")
	ErrSchedActive = errors.New("can't register schedule on active clock")
)
